// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package home

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ctaSection() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section><div class=\"max-w-6xl mx-auto px-4 sm:px-6\"><div class=\"relative px-8 py-12 md:py-20 rounded-[3rem] overflow-hidden\"><!-- Radial gradient --><div class=\"absolute flex items-center justify-center top-0 -translate-y-1/2 left-1/2 -translate-x-1/2 pointer-events-none -z-10 w-1/3 aspect-square\" aria-hidden=\"true\"><div class=\"absolute inset-0 translate-z-0 bg-purple-500 rounded-full blur-[120px] opacity-70\"></div><div class=\"absolute w-1/4 h-1/4 translate-z-0 bg-purple-400 rounded-full blur-[40px]\"></div></div><!-- Blurred shape --><div class=\"absolute bottom-0 translate-y-1/2 left-0 blur-2xl opacity-50 pointer-events-none -z-10\" aria-hidden=\"true\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"434\" height=\"427\"><defs><linearGradient id=\"bs5-a\" x1=\"19.609%\" x2=\"50%\" y1=\"14.544%\" y2=\"100%\"><stop offset=\"0%\" stop-color=\"#A855F7\"></stop> <stop offset=\"100%\" stop-color=\"#6366F1\" stop-opacity=\"0\"></stop></linearGradient></defs> <path fill=\"url(#bs5-a)\" fill-rule=\"evenodd\" d=\"m0 0 461 369-284 58z\" transform=\"matrix(1 0 0 -1 0 427)\"></path></svg></div><!-- Content --><div class=\"max-w-3xl mx-auto text-center\"><div><div class=\"inline-flex font-medium bg-clip-text text-transparent bg-gradient-to-r from-purple-500 to-purple-200 pb-3\">The security first platform</div></div><h2 class=\"h2 bg-clip-text text-transparent bg-gradient-to-r from-slate-200/60 via-slate-200 to-slate-200/60 pb-4\">Take control of your business</h2><p class=\"text-lg text-slate-400 mb-8\">All the lorem ipsum generators on the Internet tend to repeat predefined chunks as necessary, making this the first true generator on the Internet.</p><div><a class=\"btn text-slate-900 bg-gradient-to-r from-white/80 via-white to-white/80 hover:bg-white transition duration-150 ease-in-out group\" href=\"#0\">Get Started <span class=\"tracking-normal text-purple-500 group-hover:translate-x-0.5 transition-transform duration-150 ease-in-out ml-1\">-&gt;</span></a></div></div></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
