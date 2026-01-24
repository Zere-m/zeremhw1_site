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
