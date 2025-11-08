import { Card, CardContent, Typography } from "@mui/material";

const Dashboard = () => (
  <Card>
    <CardContent>
      <Typography variant="h5" gutterBottom>
        Selamat datang di Panel Admin ✨
      </Typography>
      <Typography variant="body1">
        Gunakan menu di sebelah kiri untuk mengelola artikel dan komentar.
      </Typography>
    </CardContent>
  </Card>
);

export default Dashboard;
