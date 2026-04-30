import os
import glob
import re

VUE_VIEWS_DIR = '/home/anuruprkris/Project/TASM/tasm-frontend/src/views'

files = glob.glob(os.path.join(VUE_VIEWS_DIR, '*.vue'))
attrs_to_fix = ['disabled', 'checked', 'selected', 'required', 'readonly', 'multiple', 'autofocus']

for file in files:
    with open(file, 'r', encoding='utf-8') as f:
        content = f.read()
        
    for attr in attrs_to_fix:
        content = re.sub(rf'\b{attr}=""', attr, content)
        
    with open(file, 'w', encoding='utf-8') as f:
        f.write(content)

print(f"Fixed attributes in {len(files)} files!")
