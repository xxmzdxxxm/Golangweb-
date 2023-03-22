{{define "demo"}}
昵称： {{.Name}},
{{- if .IsWin}}
恭喜，大吉大利，今晚吃鸡！
{{- else}}
遗憾，鸡被吃光了！
{{- end}}
{{- end}}