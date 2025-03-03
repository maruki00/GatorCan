import React from "react";
import {
  Drawer,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Box,
} from "@mui/material";
import { useNavigate } from "react-router-dom";
import AccountCircleIcon from "@mui/icons-material/AccountCircle";
import DashboardCustomizeRoundedIcon from "@mui/icons-material/DashboardCustomizeRounded";
import CollectionsBookmarkRoundedIcon from "@mui/icons-material/CollectionsBookmarkRounded";
import CalendarMonthRoundedIcon from "@mui/icons-material/CalendarMonthRounded";
import MailOutlineRoundedIcon from "@mui/icons-material/MailOutlineRounded";
import LogoutRoundedIcon from "@mui/icons-material/Logout";

function MyListItem({ icon, name, path, handleNavigate }) {
  return (
    <ListItem
      button
      sx={{ flexDirection: "column", alignItems: "center" }}
      onClick = {() => handleNavigate(path)}
    >
      <ListItemIcon
        sx={{
          minWidth: "unset",
          display: "flex",
          justifyContent: "center",
        }}
      >
        {icon}
      </ListItemIcon>
      <ListItemText
        primary={name}
        sx={{
          "& .MuiTypography-root": { fontSize: "0.7em" },
        }}
      />
    </ListItem>
  );
}

function StudentNavbar() {

    const handleLogout = () => {
      localStorage.clear();
      navigate("/login", { replace: true });
    };

    const navigate = useNavigate();

    const handleNavigate = (path) => {
      navigate(path, { replace: false });
    };

  return (
    <Drawer
      variant="permanent"
      anchor="left"
      sx={{
        width: 120,
        flexShrink: 0,
        "& .MuiDrawer-paper": {
          width: 100,
          boxSizing: "border-box",
          display: "flex",
          flexDirection: "column",
          justifyContent: "space-between", // Pushes logout to the bottom
        },
      }}
      PaperProps={{
        sx: {
          backgroundColor: "rgb(29, 74, 124)",
          color: "white",
        },
      }}
    >
      {/* Top Menu Items */}
      <Box sx={{ flexGrow: 1 }}>
        <List>
          <ListItem button>
            <ListItemText primary="GatorCan" />
          </ListItem>
          <MyListItem
            icon={<AccountCircleIcon sx={{ fontSize: 40, color: "white" }} />}
            name="Profile"
            path="/student-profile"
            handleNavigate={handleNavigate}
          />
          <MyListItem
            icon={
              <DashboardCustomizeRoundedIcon
                sx={{ fontSize: 40, color: "white" }}
              />
            }
            name="Dashboard"
            path="/student-dashboard"
            handleNavigate={handleNavigate}
          />
          <MyListItem
            icon={
              <CollectionsBookmarkRoundedIcon
                sx={{ fontSize: 40, color: "white" }}
              />
            }
            name="Courses"
            path="/student-courses"
            handleNavigate={handleNavigate}
          />
          <MyListItem
            icon={
              <CalendarMonthRoundedIcon sx={{ fontSize: 40, color: "white" }} />
            }
            name="Calendar"
            path="/student-calendar"
            handleNavigate={handleNavigate}
          />
          <MyListItem
            icon={
              <MailOutlineRoundedIcon sx={{ fontSize: 40, color: "white" }} />
            }
            name="Inbox"
            path="/student-inbox"
            handleNavigate={handleNavigate}
          />
        </List>
      </Box>

      {/* Logout Button at the Bottom */}
      <List>
        <ListItem
          button
          sx={{ flexDirection: "column", alignItems: "center" }}
          onClick={handleLogout} // Handle click event
        >
          <ListItemIcon
            sx={{
              minWidth: "unset",
              display: "flex",
              justifyContent: "center",
            }}
          >
            <LogoutRoundedIcon sx={{ fontSize: 40, color: "white" }} />
          </ListItemIcon>
          <ListItemText primary="Logout" />
        </ListItem>
      </List>
    </Drawer>
  );
}

export default StudentNavbar;
