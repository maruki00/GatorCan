 import { Navigate } from "react-router-dom";

 const ProtectedDashboard = ({ allowedRoles, children }) => {
   const username = localStorage.getItem("username");
   const roles = JSON.parse(localStorage.getItem("roles")) || [];

   if (!username) {
     return <Navigate to="/login" replace />;
   }
   
   if (!roles.some((role) => allowedRoles.includes(role))) {
     return <Navigate to="/dashboard" replace />;
   }

   return children;
 };

 export default ProtectedDashboard;
