import os
import glob
from bs4 import BeautifulSoup
import hashlib

directory = '/home/anuruprkris/Project/TASM/stitch_technopark_kerala_asset_management_system_v3'
files = glob.glob(os.path.join(directory, '*/code.html'))

print(f"Total screens found: {len(files)}")

tailwind_configs = {}
headers = {}
navbars = {}

for filepath in files:
    with open(filepath, 'r') as f:
        content = f.read()
        soup = BeautifulSoup(content, 'html.parser')
        
        # Check tailwind config
        tailwind_script = soup.find('script', id='tailwind-config')
        if tailwind_script:
            config_text = tailwind_script.string.strip() if tailwind_script.string else ""
            h = hashlib.md5(config_text.encode('utf-8')).hexdigest()
            tailwind_configs[h] = tailwind_configs.get(h, 0) + 1
            
        # Check header
        header = soup.find('header')
        if header:
            header_text = str(header)
            h = hashlib.md5(header_text.encode('utf-8')).hexdigest()
            headers[h] = headers.get(h, 0) + 1
            
        # Check nav (sidebar)
        nav = soup.find('nav')
        if not nav:
            nav = soup.find('aside') # sometimes it's aside
        if nav:
            nav_text = str(nav)
            h = hashlib.md5(nav_text.encode('utf-8')).hexdigest()
            navbars[h] = navbars.get(h, 0) + 1

print("\n--- Duplication Report ---")
print(f"Unique Tailwind Configs found: {len(tailwind_configs)} (Most common used in {max(tailwind_configs.values()) if tailwind_configs else 0} files)")
print(f"Unique Top Headers found: {len(headers)} (Most common used in {max(headers.values()) if headers else 0} files)")
print(f"Unique Sidebars found: {len(navbars)} (Most common used in {max(navbars.values()) if navbars else 0} files)")

