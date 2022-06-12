html_i='../index.html'
html_l='../log.html'
html_t='tmpl.html'
css='web/style.css'
src1='index.sfl'
src2='log.sfl'
sfl -nts $html_t $css $src1 $html_i && echo "LOG : compiled file to html."
sfl -nts $html_t $css $src2 $html_l && echo "LOG : compiled file to html."
