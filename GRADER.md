## System credentials:
# Super Admin 
  * Sees all data, can manage users and can comment
  * Username: admin
  * Password: password123
# Performance Analyst  
  * Sees only performance data and can comment
  * Username: performance_analyst
  * Password: password123
# Behavior Analyst
  * Sees only behavior data and can comment
  * Username: behavior_analyst
  * Password: password123
# Viewer
  * Only can view
  * Username: viewer
  * Password: password123

## Test scenario to try: 
1. Log in as admin, you should see all three reports and comments.
2. While logged in as admin, add a new test user to the database to test the user management capabilities
3. Log out
4. Log back in as any of the analysts (performance_analyst or behavior_analyst)
5. Type a test comment
6. Comfirm that you only see the allowed report for the role
7. Log out
8. Log back in as the viewer and notice that manage users is gone and the analyst comment form is hidden
9. Click the "Export to PDF" button while on the main dashboard to test the downloading feature. 
10.  While logged in as viewer, add this to the URL bar: "/manage_users.php". This should redirect you to a 403 page. You can try the same logged in as any of the analysts later

## Flaws: 
* I used the 'html2pdf.js' library for the exporting feature, which seems might get tricky if the browser window is narrow, resulting in a squished pdf. This is because the library takes a screenshot of the dom
* In my database, the table that stores logn credentials has passwords stored as plain text to make testing simpler, but this should be hashed or protected better in a real case.
* The admin can mix and match roles and visibility when creating a new user. For example, it is possible to create a viewer and assign an analyst-level permission role. In a perfect scenario, backend validation would be added.
