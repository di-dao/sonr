// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package home

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func faqsSection() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"relative\"><!-- Blurred shape --><div class=\"absolute top-0 -translate-y-1/3 left-1/2 -translate-x-1/2 ml-24 blur-2xl opacity-50 pointer-events-none -z-10\" aria-hidden=\"true\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"434\" height=\"427\"><defs><linearGradient id=\"bs3-a\" x1=\"19.609%\" x2=\"50%\" y1=\"14.544%\" y2=\"100%\"><stop offset=\"0%\" stop-color=\"#6366F1\"></stop> <stop offset=\"100%\" stop-color=\"#6366F1\" stop-opacity=\"0\"></stop></linearGradient></defs> <path fill=\"url(#bs3-a)\" fill-rule=\"evenodd\" d=\"m410 0 461 369-284 58z\" transform=\"matrix(1 0 0 -1 -410 427)\"></path></svg></div><div class=\"max-w-6xl mx-auto px-4 sm:px-6\"><div class=\"py-12 md:py-20 border-t [border-image:linear-gradient(to_right,transparent,theme(colors.slate.800),transparent)1]\"><!-- Section header --><div class=\"max-w-3xl mx-auto text-center pb-12 md:pb-20\"><div><div class=\"inline-flex font-medium bg-clip-text text-transparent bg-gradient-to-r from-purple-500 to-purple-200 pb-3\">Getting started with Sonr</div></div><h2 class=\"h2 bg-clip-text text-transparent bg-gradient-to-r from-slate-200/60 via-slate-200 to-slate-200/60 pb-4\">Everything you need to know</h2></div><!-- Columns --><div class=\"md:flex md:space-x-12 space-y-8 md:space-y-0\"><!-- Column --><div class=\"w-full md:w-1/2 space-y-8\"><!-- Item --><div class=\"space-y-2\"><h4 class=\"font-semibold\">What is Sonr?</h4><p class=\"text-slate-400\">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Quis enim lobortis scelerisque fermentum</p></div><!-- Item --><div class=\"space-y-2\"><h4 class=\"font-semibold\">What's an affordable alternative to Sonr?</h4><p class=\"text-slate-400\">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Quis enim lobortis scelerisque fermentum.</p></div><!-- Item --><div class=\"space-y-2\"><h4 class=\"font-semibold\">Can I remove the 'Powered by Sonr' branding?</h4><p class=\"text-slate-400\">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Quis enim lobortis scelerisque fermentum.</p></div></div><!-- Column --><div class=\"w-full md:w-1/2 space-y-8\"><!-- Item --><div class=\"space-y-2\"><h4 class=\"font-semibold\">What kind of data can I collect from my customers?</h4><p class=\"text-slate-400\">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Quis enim lobortis scelerisque fermentum.</p></div><!-- Item --><div class=\"space-y-2\"><h4 class=\"font-semibold\">Can I use Sonr for free?</h4><p class=\"text-slate-400\">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Quis enim lobortis scelerisque fermentum.</p></div><!-- Item --><div class=\"space-y-2\"><h4 class=\"font-semibold\">Is Sonr affordable for small businesses?</h4><p class=\"text-slate-400\">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Quis enim lobortis scelerisque fermentum.</p></div></div></div></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
