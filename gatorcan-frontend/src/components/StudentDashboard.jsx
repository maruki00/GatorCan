import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import CampaignIcon from '@mui/icons-material/Campaign';
import Typography from "@mui/material/Typography";
import EditNoteIcon from "@mui/icons-material/EditNote";
import MarkUnreadChatAltIcon from "@mui/icons-material/MarkUnreadChatAlt";
import FolderCopyIcon from "@mui/icons-material/FolderCopy";
import { Container } from "@mui/material";
import StudentNavbar from "./StudentNavbar";

import React from 'react'

function MediaCard({text1, text2, color}) {
  return (
    <Card sx={{ maxWidth: 250, margin: "20px" }} elevation={10}>
      <Container
        sx={{ backgroundColor: color, height: "150px" }}
      ></Container>
      <CardContent>
        <Typography gutterBottom variant="h8" component="div" color={color}>
          {text1}
        </Typography>
        <Typography variant="body2" sx={{ color: "text.secondary" }}>
          {text2}
        </Typography>
      </CardContent>
    </Card>
  );
}

function StudentDashboard() {

  const courses = [
    [
      "CAP5771 - Intro to Data Science",
      "CAP5771 - Intro to Data Science CAP5771 Spring 2025",
      "forestgreen",
    ],
    [
      "CEN5035 - Software Engineering",
      "CEN5035 - Software Engineering CEN5035 Spring 2025",
      "darkorchid",
    ],
    [
      "COP5556 - Program Language Principles",
      "COP5556 - Program Language Principles COP5556 Spring 2025",
      "MediumVioletRed",
    ],
  ];

  return (
    <>
      <StudentNavbar />
      <div style={{ marginLeft: "120px" }}>
        <h1>Dashboard</h1>
        <hr />
        <br></br>
        <h4>My Courses</h4>
        <div
          style={{
            display: "flex",
            flexWrap: "wrap",
            justifyContent: "left",
            flexDirection: "row",
            width: "80%",
          }}
        >
          {courses.map((course, index) => {
            return (
              <MediaCard
                key={index}
                text1={course[0]}
                text2={course[1]}
                color={course[2]}
              ></MediaCard>
            );
          })}
        </div>
      </div>
    </>
  );
}

export default StudentDashboard