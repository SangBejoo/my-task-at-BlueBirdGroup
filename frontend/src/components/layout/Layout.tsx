import React from 'react';
import { Box, Drawer, AppBar, Toolbar, Typography, List, ListItem, ListItemIcon, ListItemText, ListItemButton } from '@mui/material';
import DashboardIcon from '@mui/icons-material/Dashboard';
import MapIcon from '@mui/icons-material/Map';
import MonitorIcon from '@mui/icons-material/Monitor';
import { useNavigate } from 'react-router-dom';

const DRAWER_WIDTH = 240;

const Layout = ({ children }: { children: React.ReactNode }) => {
  const navigate = useNavigate();

  const menuItems = [
    { text: 'Dashboard', icon: <DashboardIcon />, path: '/' },
    { text: 'Parking Map', icon: <MapIcon />, path: '/map' },
    { text: 'Monitoring', icon: <MonitorIcon />, path: '/monitoring' },
  ];

  return (
    <Box sx={{ display: 'flex' }}>
      <AppBar position="fixed">
        <Toolbar>
          <Typography variant="h6">Parking Space Monitor</Typography>
        </Toolbar>
      </AppBar>
      
      <Drawer
        variant="permanent"
        sx={{
          width: DRAWER_WIDTH,
          '& .MuiDrawer-paper': { width: DRAWER_WIDTH, boxSizing: 'border-box' },
        }}
      >
        <Toolbar />
        <List>
          {menuItems.map((item) => (
            <ListItem key={item.text} disablePadding>
              <ListItemButton onClick={() => navigate(item.path)}>
                <ListItemIcon>{item.icon}</ListItemIcon>
                <ListItemText primary={item.text} />
              </ListItemButton>
            </ListItem>
          ))}
        </List>
      </Drawer>

      <Box
        component="main"
        sx={{ flexGrow: 1, p: 3, width: `calc(100% - ${DRAWER_WIDTH}px)` }}
      >
        <Toolbar />
        {children}
      </Box>
    </Box>
  );
};

export default Layout;