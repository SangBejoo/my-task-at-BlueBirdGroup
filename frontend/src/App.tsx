import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import React, { Suspense } from 'react';
import { CircularProgress } from '@mui/material';
import Layout from './components/layout/Layout';
import Dashboard from './pages/Dashboard';
import ParkingMap from './pages/ParkingMap';
import Monitoring from './pages/Monitoring';

const LoadingFallback = () => (
  <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
    <CircularProgress />
  </div>
);

function App() {
  return (
    <BrowserRouter>
      <Suspense fallback={<LoadingFallback />}>
        <Layout>
          <Routes>
            <Route path="/" element={<Dashboard />} />
            <Route path="/map" element={<ParkingMap />} />
            <Route path="/monitoring" element={<Monitoring />} />
            <Route path="*" element={<Navigate to="/" replace />} />
          </Routes>
        </Layout>
      </Suspense>
    </BrowserRouter>
  );
}

export default App;