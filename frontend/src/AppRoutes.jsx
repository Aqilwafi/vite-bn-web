import { Routes, Route } from 'react-router-dom';
import Home from './pages/HomePage';
import PostDetail from "./pages/PostDetail";


export default function AppRoutes() {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/post/:slug" element={<PostDetail />} />
    </Routes>
  );
}
