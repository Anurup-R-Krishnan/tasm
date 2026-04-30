import json
import re
from bs4 import BeautifulSoup

with open('/home/anuruprkris/Project/TASM/stitch_technopark_kerala_asset_management_system_v3/consumables_supplies_tracker/code.html', 'r') as f:
    soup = BeautifulSoup(f.read(), 'html.parser')

script = soup.find('script', id='tailwind-config').string

# Extract the JSON part
json_str = script.replace('tailwind.config =', '').strip()
if json_str.endswith(';'):
    json_str = json_str[:-1]

config = json.loads(json_str)

theme = config.get('theme', {}).get('extend', {})
colors = theme.get('colors', {})
fonts = theme.get('fontFamily', {})

out_css = "@import 'tailwindcss';\n\n@theme {\n"

for name, val in colors.items():
    if isinstance(val, str):
        out_css += f"  --color-{name}: {val};\n"

for name, val in fonts.items():
    if isinstance(val, list):
        out_css += f"  --font-{name}: '{val[0]}', sans-serif;\n"

out_css += "}\n\n"
out_css += "body {\n  @apply bg-canvas text-body font-body text-text-primary antialiased min-h-screen;\n}\n"

with open('/home/anuruprkris/Project/TASM/tasm-frontend/src/style.css', 'w') as f:
    f.write(out_css)
