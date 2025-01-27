// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package inquiryform

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func InquiryForm() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"inquiry-form\" class=\"fixed z-10 inset-0 overflow-y-auto\"><div class=\"flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0\"><div class=\"fixed inset-0 transition-opacity\" aria-hidden=\"true\"><div class=\"absolute inset-0 bg-gray-500 opacity-75\"></div></div><span class=\"hidden sm:inline-block sm:align-middle sm:h-screen\" aria-hidden=\"true\">&#8203;</span><!-- add the close button here --><div class=\"inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full\"><button class=\"float-right flex items-center justify-center w-10 h-10 bg-gray-200 text-gray-600 rounded-full hover:bg-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-400 m-2\" hx-delete=\"/inquiry\" hx-trigger=\"click\" hx-swap=\"outerHTML\" hx-target=\"#inquiry-form\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"w-6 h-6\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M6 18L18 6M6 6l12 12\"></path></svg></button><div class=\"bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4\"><div><form hx-post=\"/inquiry\" hx-target=\"#inquiry-form\" hx-swap=\"outerHTML\" method=\"post\"><div class=\"mb-4\"><label for=\"name\" class=\"block text-gray-700 text-sm font-bold mb-2\">Vardas:</label> <input type=\"text\" name=\"name\" class=\"shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline\" required></div><div class=\"mb-4\"><label for=\"email\" class=\"block text-gray-700 text-sm font-bold mb-2\">El.Paštas:</label> <input type=\"email\" name=\"email\" class=\"shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline\" required></div><div class=\"mb-4\"><label for=\"description\" class=\"block text-gray-700 text-sm font-bold mb-2\">Užklausa</label> <textarea name=\"description\" rows=\"4\" class=\"shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline\" required></textarea></div><div class=\"bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse\"><button type=\"submit\" class=\"w-full inline-flex justify-center rounded-md border border-transparent\n\t\t\t\t\t\t\t\t\tshadow-sm px-4 py-2 \n\t\t\t\t\t\t\t\t\tbg-yellow-600 text-base font-medium\n\t\t\t\t\t\t\t\t\ttext-white hover:bg-yellow-700 focus:outline-none\n\t\t\t\t\t\t\t\t\tfocus:ring-2 focus:ring-offset-2 focus:ring-yellow-500 \n\t\t\t\t\t\t\t\t\tsm:ml-3 sm:w-auto sm:text-sm\">Siųsti!</button></div></form></div></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
