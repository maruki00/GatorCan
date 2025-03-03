import { Routes, Route } from "react-router-dom";

import Login from "./components/Login";
import ProtectedRoute from "./components/ProtectedRoute";
import ProtectedDashboard from "./components/ProtectedDashboard";
import AdminDashboard from "./components/AdminDashboard";
import StudentDashboard from "./components/StudentDashboard";
import InstructorDashboard from "./components/InstructorDashboard";
import UserRegistration from "./components/UserRegistration";
import StudentCalendar from "./components/StudentCalendar";
import StudentInbox from "./components/StudentInbox";
import StudentProfile from "./components/StudentProfile";
import StudentCourses from "./components/StudentCourses";

import "./App.css";

function App() {
  return (
    <Routes>
      <Route path="login" element={<Login />} />
      <Route path="dashboard" element={<ProtectedRoute />} />

      {/* Protecting dashboard routes */}
      <Route
        path="admin-dashboard"
        element={
          <ProtectedDashboard allowedRoles={["admin"]}>
            <AdminDashboard />
          </ProtectedDashboard>
        }
      />
      <Route
        path="student-dashboard"
        element={
          <ProtectedDashboard allowedRoles={["student"]}>
            <StudentDashboard />
          </ProtectedDashboard>
        }
      />
      <Route
        path="instructor-dashboard"
        element={
          <ProtectedDashboard allowedRoles={["instructor"]}>
            <InstructorDashboard />
          </ProtectedDashboard>
        }
      />

      <Route
        path="user-registration"
        element={
          <ProtectedDashboard allowedRoles={["admin"]}>
            <UserRegistration />
          </ProtectedDashboard>
        }
      />

      <Route
        path="student-profile"
        element={
          <ProtectedDashboard allowedRoles={["student"]}>
            <StudentProfile />
          </ProtectedDashboard>
        }
      />

      <Route
        path="student-calendar"
        element={
          <ProtectedDashboard allowedRoles={["student"]}>
            <StudentCalendar />
          </ProtectedDashboard>
        }
      />

      <Route
        path="student-inbox"
        element={
          <ProtectedDashboard allowedRoles={["student"]}>
            <StudentInbox />
          </ProtectedDashboard>
        }
      />

      <Route
        path="student-courses"
        element={
          <ProtectedDashboard allowedRoles={["student"]}>
            <StudentCourses />
          </ProtectedDashboard>
        }
      />

      <Route path="*" element={<Login />} />
    </Routes>
  );
}

export default App;
