# Sprint 1 Report

### Project Name: Gator-Can 

### Project Type : E-Learning-Platform

## Visual Demo Links
- [Backend](https://drive.google.com/file/d/1izfWzIRqQzYru6UtwQtQb5mZ8EN1Hd3_/view?usp=sharing)
- [Frontend](https://drive.google.com/file/d/1oA1x6nKK5xaCnk7OiqgrPVOJlzpUx9mM/view?usp=sharing)

## Overview of Sprint 1 :

We established the core functionality of the E-Learning-Platform during sprint 1. Our main goal was to design the infrastructure and lay the groundwork that we would iteratively improve in future sprints.

## What was accomplished :

### Frontend using React : Done by Harsh and Navnit
We designed and developed a static interface using React and JavaScript.

The login page as well as user dashboard was created by the frontend team and was designed to be adaptable to future feature enhancements.

Specific time was allocated to discuss how the layout could be better modified to be intuitive and responsive.

### Backend using Go : Done by Muthukumaran and Wael
The backend was built using Golang as it provided an efficient server framework. 

The team focused on setting up the core server architecture and ensured that the frontend could be connected easily connected.

For Sprint 1 - the backend is mostly static but we laid the groundwork for future dynamic content like course management or user registration and it was designed to require minimal changes in the future when we add more features by making it compatible to improvements.

Password Encryption:
We prioritized user security and privacy and used encryption for passwords rather than storing them as plaintext.
To protect user data, we implemented password encryption using the bcrypt library for hashing the passwords before storing them in the database. 
This ensures that user passwords are stored securely and reduces the threat of unauthorized access.

Role-Based Access Control :
We implemented role-based access control to manage different user types on the platform - including Admins, Instructors, and Students.
This makes it so users possess an approporiate level of clearance and cannot access files they are not authorized to view.
Admins can manage the entire platform whereas instructors can manage individual courses and students have access to learning content.

## User Stories

- Allow multiple types of users, such as Students and Admins.
- Provide a single login page for all users.
- Display appropriate login messages to users.
- Persist user login sessions even after a page reload.
- After logging in, show dashboards specific to the user type.
- The Student Dashboard will have a navigation bar for different sections, such as My Courses and My Profile, with an attractive UI.
- The Student Dashboard will include a logout button to allow users to log out.
- The Student Dashboard will display a list of courses the student has enrolled in.
- The Admin Dashboard will have a clean UI and a logout button.
- The Admin Dashboard will display a list of available actions on the homepage.
- Clicking on any action will redirect the user to the corresponding page.
- Clicking "Add User" will navigate the admin to the Add User page.
- After filling out the Add User form, validate the inputs and display success or error messages accordingly.

## Issues Resolved

### Backend

- Implement Fetch User Details API: Develop an API that fetches user details based on authentication tokens.
- Implement JWT Authentication: JWT-based authentication to be set up, thereby allowing users to securely log in and maintain sessions.
- Define user roles (Admin, Instructor, Student, Teaching Assistant): Establish user roles as well as associated permissions within the system to ensure access control.
- Implement Role-Based Access Control (RBAC): Implement a system where different user roles have specific permissions and access levels.
- Implement authorization checks in middleware: Develop middleware logic to enforce role-based access control throughout the application.
- Restrict registration to Admin only: Limit user registration so that only Admins can add new users.
- Write unit tests for role-based access: Create and run unit tests to verify that RBAC was working correctly.
- Create /register endpoint: Develop the user registration endpoint, allowing Admins to create new accounts.
- Validate input and store user details securely: Ensure all user-provided data was validated and securely stored in the database.
- Write unit tests for user registration: Develop unit tests specifically for the user registration process to ensure data integrity and security.
- Create /users/me endpoint: Creat an endpoint that returns the authenticated user"s details, allowing them to access their profile.
- Fetch user details from the database: Implement functionality to retrieve user details from the database for profile and role management.
- Create /login endpoint: Develop the login API endpoint to handle user authentication and return a JWT token.
- Validate user credentials and generate JWT token: Implement logic to validate login credentials and generate a secure JWT token for authenticated users.
- Implement middleware for token verification: Develop a middleware to verify JWT tokens, ensuring only authenticated users can access protected routes.
- Write unit tests: Conduct unit testing across backend functionalities to identify and resolve potential issues.
- Write unit tests for authentication flow: Create and execute unit tests to ensure the authentication process works as expected, covering login and token generation.

### Frontend

- Initialize React app with Material-UI: Set up a new React project using Vite with TypeScript and install Material-UI (MUI) to build a modern user interface. Make sure all necessary dependencies are installed and configure MUI’s theme for consistent styling across the app.
- Set up project folder structure: Organize the project files into separate folders for components, pages, services, hooks, and utilities. This will make the project easier to manage and scale as new features are added.
- Design Login page UI: Build the login page using Material-UI components. Include input fields for email and password, a submit button, and space for error messages. Ensure the layout looks clean and is easy to use.
- Implement form validation (email & password): Add validation to check if the email is in the correct format and the password meets basic security requirements. Show error messages if the user enters invalid input and prevent form submission until all fields are valid.
- Write login page UI tests using Jest: Use Jest and React Testing Library to write tests for the login form. Check if the form renders correctly, validation errors show up when needed, and the submit button is disabled for invalid input.
- Create API service for authentication: Set up a separate service file to handle authentication API requests. This will include functions for login and logout, making it easier to manage API calls in one place.
- Store JWT token in local storage: After a successful login, store the received JWT token in local storage so users stay logged in even after refreshing the page. Ensure secure handling of the token.
- Redirect users based on roles: After logging in, check the user’s role and redirect them to the appropriate dashboard (e.g., Admins go to the admin panel, Students to the student dashboard).
- Write unit tests for API calls: Use Jest to test API service functions. Ensure that API calls return expected results and handle errors properly. Mock API responses to simulate different scenarios.
- Integrate Login API with React UI: Connect the login form with the authentication API. When users enter their credentials and submit the form, send the data to the backend and handle the response accordingly. Show error messages for failed logins.
- Implement User Registration Form (Admin Only): Create a registration form that only Admins can access. This form will allow Admins to add new users by providing their email, password, and role.
- Design user registration form: Build the UI for the registration form using Material-UI components. Include fields for email, password, and role selection.
- Validate inputs (role selection, email, password): Add validation to ensure the email is correctly formatted, the password meets security requirements, and a role is selected before submitting the form.
- Call /register API on form submission: When the form is submitted, send the data to the /register API endpoint to create a new user. Show appropriate messages based on the response.
- Show success/error messages: Display a success message if registration is successful and an error message if something goes wrong (e.g., invalid input, API failure).
- Write register page UI tests: Use Jest and React Testing Library to test the registration form. Ensure form elements render correctly, validation works as expected, and API calls behave properly in different scenarios.
- Implement role-based dashboard rendering (Admin vs. Student vs. Instructor): Modify the UI to show different dashboard components based on the user’s role. Ensure that Admins see admin-only features, while Students and Instructors have access to their respective sections.
- Hide admin features for non-admin users: Restrict visibility of admin-only features so that non-admin users cannot access or interact with them.
- Test UI access based on role: Write tests to verify that the correct dashboard elements appear based on the logged-in user’s role. Ensure unauthorized users cannot access restricted features.

### General Issues

- Discussion -> Branching Strategies: Held discussions on the best branching strategy to follow for efficient collaboration and code management.

## Issues completed:

We have completed all the issues listed above

## Issues remaining:

No issues are remaining from the above list of issues
