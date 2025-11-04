
import { CssBaseline, Box } from '@mui/material';
import { Route, Routes } from 'react-router-dom';
import NavBar from './features/navbar/NavBar';
import AboutPage from './pages/about/AboutPage';
import HomePage from './pages/home/HomePage';

function App() {
  return (
    <div>
      <CssBaseline />
      <NavBar />
      <Box sx={{ padding: 2, height: 'calc(100vh - 64px)' }}>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/about" element={<AboutPage />} />
        </Routes>
      </Box>

    </div>
  );
}

export default App;
