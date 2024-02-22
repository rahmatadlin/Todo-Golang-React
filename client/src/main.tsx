import React from 'react';
import ReactDOM from 'react-dom/client';
import { MantineProvider } from '@mantine/core'; // Import MantineProvider
import App from './App';
import './index.css';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    {/* Wrap your app with MantineProvider */}
    <MantineProvider>
      <App />
    </MantineProvider>
  </React.StrictMode>
);
