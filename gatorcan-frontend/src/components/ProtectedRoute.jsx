import { Navigate, useNavigate } from "react-router-dom";

const ProtectedRoute = () => {
  const username = localStorage.getItem("username");
  const roles = JSON.parse(localStorage.getItem("roles")) || [];

  if (!username) {
    return <Navigate to="/login" replace />;
  }

  if (roles.length === 1) {
    if (roles.includes("admin"))
      return <Navigate to="/admin-dashboard" replace />;
    if (roles.includes("student"))
      return <Navigate to="/student-dashboard" replace />;
    if (roles.includes("instructor"))
      return <Navigate to="/instructor-dashboard" replace />;
  }

  return <RoleSelection roles={roles} />;
};

const RoleSelection = ({ roles }) => {
  const navigate = useNavigate();

  return (
    <div>
      <h2>Login as:</h2>
      {roles.map((role) => (
        <button key={role} onClick={() => navigate(`/${role}-dashboard`)}>
          {role.charAt(0).toUpperCase() + role.slice(1)}
        </button>
      ))}
    </div>
  );
};

export default ProtectedRoute;
