package badge

var flatTemplate = stripXmlWhitespace(`
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{{.Bounds.Dx}}" height="20">
  <title>{{.Subject | html}}: {{.Status | html}}</title>
  <linearGradient id="s" x2="0" y2="100%">
    <stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
    <stop offset="1" stop-opacity=".1"/>
  </linearGradient>
  <clipPath id="r">
    <rect width="{{.Bounds.Dx}}" height="20" rx="3" fill="#fff" />
  </clipPath>
  <g mask="clip-path(#r)">
    <rect width="{{.Bounds.SubjectDx}}" height="20" fill="#555"/>
    <rect x="{{.Bounds.SubjectDx}}" width="{{.Bounds.StatusDx}}" height="20" fill="{{or .Color "#4c1" | html}}"/>
    <rect width="{{.Bounds.Dx}}" height="20" fill="url(#s)"/>
  </g>
  <g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="110">
    <text aria-hidden="true" x="{{.Bounds.SubjectX}}" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)">{{.Subject | html}}</text>
    <text x="{{.Bounds.SubjectX}}" y="140" transform="scale(.1)">{{.Subject | html}}</text>
    <text aria-hidden="true" x="{{.Bounds.StatusX}}" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)">{{.Status | html}}</text>
    <text x="{{.Bounds.StatusX}}" y="140" transform="scale(.1)" fill="#fff">{{.Status | html}}</text>
  </g>
</svg>
`)
