// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package home

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/didao-org/sonr/internal/components/base"
	"github.com/didao-org/sonr/internal/components/views/landing"
	"github.com/labstack/echo/v4"
)

func Page(ctx echo.Context) templ.Component {
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
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			templ_7745c5c3_Err = landing.Navbar().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <main class=\"grow\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = heroSection().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = logosSection().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = introFeaturesSection().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = primaryFeaturesSection().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = secondaryFeaturesSection().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = faqsSection().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = ctaSection().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = landing.Footer().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</main>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = base.PageLayout().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func logosSection() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Clients --><section><div class=\"relative max-w-6xl mx-auto px-4 sm:px-6\"><!-- Particles animation --><div class=\"absolute inset-0 max-w-6xl mx-auto px-4 sm:px-6\"><div class=\"absolute inset-0 -z-10\" aria-hidden=\"true\"><canvas data-particle-animation data-particle-quantity=\"5\"></canvas></div></div><div class=\"py-12 md:py-16\"><div class=\"overflow-hidden\"><!-- Carousel built with Swiper.js [https://swiperjs.com/] --><!-- * Initialized in src/js/main.js --><!-- * Custom styles in src/css/additional-styles/theme.scss --><div class=\"clients-carousel swiper-container relative before:absolute before:inset-0 before:w-32 before:z-10 before:pointer-events-none before:bg-gradient-to-r before:from-slate-900 after:absolute after:inset-0 after:left-auto after:w-32 after:z-10 after:pointer-events-none after:bg-gradient-to-l after:from-slate-900\"><div class=\"swiper-wrapper !ease-linear select-none items-center\"><!-- Carousel items --><div class=\"swiper-slide !w-auto\"><img src=\"https://cdn.sonr.build/images/client-01.svg\" alt=\"W3C Logo\" height=\"20\"></div><div class=\"swiper-slide !w-auto\"><img src=\"https://cdn.sonr.build/images/client-07.svg\" alt=\"Torch Capital\" width=\"230\"></div><div class=\"swiper-slide !w-auto\"><img src=\"https://cdn.sonr.build/images/client-04.svg\" alt=\"Matrix\" width=\"100\"></div><div class=\"swiper-slide !w-auto\"><img src=\"https://cdn.sonr.build/images/client-06.svg\" alt=\"Outlier Ventures\" width=\"180\"></div><div class=\"swiper-slide !w-auto\"><img src=\"https://cdn.sonr.build/images/client-05.svg\" alt=\"DIF\" width=\"100\"></div><div class=\"swiper-slide !w-auto\"><img class=\"mt-1\" src=\"https://cdn.sonr.build/images/client-03.svg\" alt=\"IBC Protocol\" width=\"160\"></div><div class=\"swiper-slide !w-auto\"><img class=\"mt-2\" src=\"https://cdn.sonr.build/images/client-09.svg\" alt=\"Fido Alliance Logo\" width=\"88\"></div><div class=\"swiper-slide !w-auto\"><img src=\"https://cdn.sonr.build/images/client-02.svg\" alt=\"Shima Capital Logo\" width=\"172\"></div></div></div></div></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}