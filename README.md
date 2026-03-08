
# Derisk checkpoint
## Dashboard URL: 
  https://reporting.zeremhw1.site/dashboard.php

## Username & password for reporting dashboard: 
  User: grader
  Password: Grader123

## Team Member: 
  Zere Mukanova

## Point 1: 
  I decided to continue with PHP for a simpler solution, and created a PHP session based login system for the dashboard. This has simple logic: if someone tries to access the dashboard by bypassing the login screen, they will be redirected back to the login
## Point 2: 
  I sent a request using the fetch() API and GET to grab the data from collector server. I also configured CORS headers on the API to allow my reporting domain to read the JSON payload. My Data table is a simple HTML table created using plain JavaScript, where all table cells are populated using textContent. This prevents XSS attacks because the tags are not revealed.
## Point 3: 
  I used Chart.js to create immediate feedback, using a simple frequency count on my JSON data (for each metric). 




# Homework 3

## Link to website: 
  https://zeremhw1.site/

## Team Member: 
  Zere Mukanova

## IP address: 
  64.23.244.69

## Grader login to server:
  Username: grader
  Password: Grader123


## Website login:
  Username: zere
  Password: CSE135winter2026

## Changes to collector.js:

Added a session creator (getSessionId): A function that generates a unique session ID using crypto.randomUUID() and saves it to the browser's sessionStorage to tie all data sent during a single visit to one user.

Added the sessionId into the reportPerf() object so the database knows exactly which session the data belongs to.


Added CSS Verification (checkCssAllowed): A function that temporarily injects a hidden <div> into the DOM to verify if the browser is successfully processing CSS rules.

Added image verification (checkImagesAllowed): A function that attempts to load a 1x1 pixel base64 image to verify if the browser allows image rendering.

initialBrowserData(): Upgraded the static data collection to include windowWidth, windowHeight, screenWidth, screenHeight, and javascriptEnabled, imagesAllowed, cssAllowed.

Replaced console.log with a POST Request: Removed the placeholder console.log(payload) inside reportPerf() and replaced it with a fetch() POST request pointing to my REST endpoint (https://collector.zeremhw1.site/api/metrics/index.php).

Added keepalive: true: to the fetch request to ensure the browser finishes sending the data to my server even if the user is in the middle of closing the tab.

Added blobal tracking state: Created activityQueue (array), idleTimer, lastActivityTime, isIdle, and idleStartTime variables at the top of the file to manage user state.

Created a queueing function (logActivity) to package custom user events and push them into the activityQueue array rather than sending them immediately.

Implemented idle time functionality: A resetIdleTimer function that marks a user as "idle" after 2 seconds of no input, and logs exactly how long they were idle when they return (idleEnd and idleStart events).

Mouse tracking: Added event listeners to track mousedown and mousemove.

Added scroll and keyboard tracking: Added event listeners to log window scroll positions and exact keystrokes.

Added page lifecycle tracking: listeners to log when the user enters the page and when they leave the page.

sendActivityData(): a dedicated function that takes everything in the activityQueue, packages it into a single JSON payload labeled as userActivity, clears the queue, and sends it to the REST endpoint.

Added interval timer: setInterval(sendActivityData, 5000) inside the initialization step to automatically fire off batches of queued user activity every 5 seconds

Added logManualLoadTime():from a custom metric to calculate the exact page load time in milliseconds by using setTimeout to wait for the load event to completely finish, then subtracting the Navigation Timing API's startTime from loadEventEnd.

Changed window.onload: Appended logManualLoadTime() and initActivityTracking() to the end of the window.onload block so all custom trackers initialize automatically as soon as the DOM is ready.





# Homework 2: 
## Link to website: 
https://zeremhw1.site/

## Approach of free-choice analytics: 
I chose Microsoft Clarity after briefly looking at other options like Plausible and Umami, but Clarity was the easiest to set up quickly and is completely free, which was the main reason. It only required adding a simple script tag to my pages, and provided helpful features like heatmaps and session replays to see how users interact with the site. 

## Team Member: 
Zere Mukanova

## IP address: 
64.23.244.69

## Grader login to server:
  Username: grader
  Password: Grader123


## Website login:
  Username: zere
  Password: CSE135winter2026
  
# Source code for programs in /public_html/cgi-bin!!




# Homework 1: 

## zeremhw1_site
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
