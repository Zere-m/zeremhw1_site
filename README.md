## Homework 2: 
# Link to website: 
https://zeremhw1.site/

# Approach of free-choice analytics: 
I chose Microsoft Clarity after briefly looking at other options like Plausible and Umami, but Clarity was the easiest to set up quickly and is completely free, which was the main reason. It only required adding a simple script tag to my pages, and provided helpful features like heatmaps and session replays to see how users interact with the site. 

# Team Member: 
Zere Mukanova

# IP address: 
64.23.244.69

# Grader login to server:
  Username: grader
  Password: Grader123


# Website login:
  Username: zere
  Password: CSE135winter2026
  
# Source code for programs in /public_html/cgi-bin!!




## Homework 2: 

# zeremhw1_site
Team members:
  Zere Mukanova

Grader login:
  Username: grader
  Password: Grader123

Website link:
  https://zeremhw1.site/

GitHub auto-deploy setup:
  I configured a post-receive webhook on the server.
  When I push changes to GitHub, the webhook triggers a script that runs git pull and updates the site files automatically.
  The website files are stored in:
  /var/www/zeremhw1.site/public_html

Website login:
  Username: zere
  Password: CSE135winter2026

Compression summary:
  After enabling mod_deflate, I checked the site in Chrome DevTools.
  I enabled compression, added a config file to compress HTML, CSS, and JS, then restarted Apache.
  After this, the response header showed Content-Encoding: gzip, and the Transferred size was smaller than the Resource size, proving the files were compressed before being sent.

Server header summary:
  I first tried using mod_headers to change the Server header, but Apache wouldn’t let me override it.
  Then I installed and enabled mod_security2 and used SecServerSignature to change the Server header.
  After restarting Apache, the header correctly showed CSE 135.
