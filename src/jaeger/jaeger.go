package jaeger

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
)

// Wrapper around jaeger.RemoteReporter that adds the environment tag to the spans
type customRemoteReporter struct {
	reporter jaeger.Reporter
}

func (r *customRemoteReporter) Report(span *jaeger.Span) {
	span.SetTag("environment", os.Getenv("ENVIRONMENT"))
	r.reporter.Report(span)
}

func (r *customRemoteReporter) Close() {
	r.reporter.Close()
}

func NewCustomRemoteReporter(sender jaeger.Transport, opts ...jaeger.ReporterOption) jaeger.Reporter {
	return &customRemoteReporter{
		reporter: jaeger.NewRemoteReporter(sender, opts...),
	}
}

// Creates a new Jaeger tracer client with UDP transport
func New(service string, logger log.Logger) (opentracing.Tracer, io.Closer) {

	enabled, err := strconv.ParseBool(os.Getenv("ENABLE_TRACING"))
	if err != nil || !enabled {
		level.Info(logger).Log(
			"message", "ENABLE_TRACING is either false or undefined, tracing will be disabled.",
			"err", err,
			"enable_tracing", enabled,
		)
		return opentracing.NoopTracer{}, ioutil.NopCloser(bytes.NewBuffer([]byte{}))
	}

	addr := os.Getenv("JAEGER_ADDRESS") + ":" + os.Getenv("JAEGER_PORT")

	transport, err := jaeger.NewUDPTransport(addr, 0)
	if err != nil {
		level.Error(logger).Log(
			"message", fmt.Sprintf("Unable to start Jaeger transport on %s, tracing will be disabled", addr),
			"err", err,
			"address", addr,
		)
		return opentracing.NoopTracer{}, ioutil.NopCloser(bytes.NewBuffer([]byte{}))
	}

	sampling, err := strconv.ParseFloat(os.Getenv("TRACING_SAMPLING_RATE"), 64)
	if err != nil {
		level.Info(logger).Log(
			"message", "TRACING_SAMPLING_RATE must be a valid float between 0 and 1.",
			"err", err,
		)
		return opentracing.NoopTracer{}, ioutil.NopCloser(bytes.NewBuffer([]byte{}))
	}

	sampler, err := jaeger.NewProbabilisticSampler(sampling)
	if err != nil {
		level.Info(logger).Log(
			"message", "Unable to create sampler, tracing will be disabled.",
			"err", err,
		)
		return opentracing.NoopTracer{}, ioutil.NopCloser(bytes.NewBuffer([]byte{}))
	}

	return jaeger.NewTracer(
		service,
		sampler,
		NewCustomRemoteReporter(transport, jaeger.ReporterOptions.Logger(jaeger.StdLogger)),
		jaeger.TracerOptions.Logger(jaeger.StdLogger))

}
