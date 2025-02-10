import * as React from "react";
import { useNavigate } from "react-router-dom";

import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";
import Menu from "@mui/material/Menu";
import Container from "@mui/material/Container";
import Avatar from "@mui/material/Avatar";
import Tooltip from "@mui/material/Tooltip";
import MenuItem from "@mui/material/MenuItem";

import Grid from "@mui/material/Grid";
import PersonAddOutlinedIcon from '@mui/icons-material/PersonAddOutlined';
import PersonRemoveOutlinedIcon from "@mui/icons-material/PersonRemoveOutlined";
import CollectionsBookmarkOutlinedIcon from "@mui/icons-material/CollectionsBookmarkOutlined";
import EditNoteOutlinedIcon from "@mui/icons-material/EditNoteOutlined";
import BorderColorOutlinedIcon from "@mui/icons-material/BorderColorOutlined";
import BackspaceOutlinedIcon from "@mui/icons-material/BackspaceOutlined";
import PersonSearchOutlinedIcon from "@mui/icons-material/PersonSearchOutlined";
import ContentPasteSearchOutlinedIcon from "@mui/icons-material/ContentPasteSearchOutlined";
import QueryStatsOutlinedIcon from "@mui/icons-material/QueryStatsOutlined";

import {
  Card,
  CardActionArea,
  CardContent,
} from "@mui/material";

const settings = ["Profile", "Logout"];

function AdminDashboard() {
  const navigate = useNavigate();

  const handleLogout = () => {
    localStorage.clear();
    navigate("/login", { replace: true });
  };

  const [anchorElUser, setAnchorElUser] = React.useState(null);

  const handleOpenUserMenu = (event) => {
    setAnchorElUser(event.currentTarget);
  };

  const handleCloseUserMenu = () => {
    setAnchorElUser(null);
  };

  const addUser = () => {
    navigate("/user-registration");
  }

  const tools = [
    [<PersonAddOutlinedIcon />, "LightSalmon", "Add User", addUser],
    [<PersonRemoveOutlinedIcon />, "LightPink", "Delete User"],
    [<BorderColorOutlinedIcon />, "Salmon", "Edit User"],
    [<CollectionsBookmarkOutlinedIcon />, "Coral", "Add Course"],
    [<BackspaceOutlinedIcon />, "Gold", "Delete Course"],
    [<EditNoteOutlinedIcon />, "Tomato", "Edit Course"],
    [<PersonSearchOutlinedIcon />, "PapayaWhip", "View Users"],
    [<ContentPasteSearchOutlinedIcon />, "Khaki", "View Courses"],
    [<QueryStatsOutlinedIcon />, "Crimson", "Statistics"],
  ];

  return (
    <div>
      <AppBar position="static">
        <Container maxWidth="xl">
          <Toolbar disableGutters>
            <Typography
              variant="h6"
              noWrap
              component="a"
              href="/"
              sx={{
                mr: 2,
                display: { xs: "none", md: "flex" },
                fontFamily: "monospace",
                fontWeight: 700,
                letterSpacing: ".2rem",
                color: "inherit",
                textDecoration: "none",
              }}
              onClick={(e) => {}}
            >
              GATORCAN-ADMIN
            </Typography>

            <Box sx={{ flexGrow: 1 }} />

            <Box sx={{ flexGrow: 0 }}>
              <Tooltip title="Open settings">
                <IconButton onClick={handleOpenUserMenu} sx={{ p: 0 }}>
                  <Avatar alt="Remy Sharp" src="/static/images/avatar/2.jpg" />
                </IconButton>
              </Tooltip>
              <Menu
                sx={{ mt: "45px" }}
                id="menu-appbar"
                anchorEl={anchorElUser}
                anchorOrigin={{
                  vertical: "top",
                  horizontal: "right",
                }}
                keepMounted
                transformOrigin={{
                  vertical: "top",
                  horizontal: "right",
                }}
                open={Boolean(anchorElUser)}
                onClose={handleCloseUserMenu}
              >
                {settings.map((setting) => (
                  <MenuItem
                    key={setting}
                    onClick={
                      setting === "Logout" ? handleLogout : handleCloseUserMenu
                    }
                  >
                    <Typography sx={{ textAlign: "center" }}>
                      {setting}
                    </Typography>
                  </MenuItem>
                ))}
              </Menu>
            </Box>
          </Toolbar>
        </Container>
      </AppBar>
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        minHeight="100vh"
      >
        {/* Fully centered layout */}
        <Grid container spacing={5} maxWidth={600}>
          {" "}
          {/* Adjusted max width to center grid tightly */}
          {tools.map((_, index) => (
            <Grid
              item
              key={index}
              xs={4}
              display="flex"
              justifyContent="center"
            >
              {" "}
              {/* Ensuring 3 cards per row */}
              <Card
                sx={{
                  width: 120,
                  height: 120,
                  backgroundColor: tools[index][1],
                }}
                onClick={tools[index][3]}
              >
                <CardActionArea sx={{ height: "100%" }}>
                  <CardContent
                    sx={{
                      display: "flex",
                      flexDirection: "column",
                      alignItems: "center",
                      justifyContent: "center",
                      height: "100%",
                    }}
                  >
                    {tools[index][0]}
                    <Typography variant="caption">{tools[index][2]}</Typography>
                  </CardContent>
                </CardActionArea>
              </Card>
            </Grid>
          ))}
        </Grid>
      </Box>
    </div>
  );
}

export default AdminDashboard;
