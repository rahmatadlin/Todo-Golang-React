import React from 'react';
import { createRoot } from 'react-dom/client';
import { MantineProvider } from '@mantine/core';
import App from './App';

const rootElement = document.getElementById('root');

if (rootElement) {
  const root = createRoot(rootElement);

  root.render(
    <React.StrictMode>
      <MantineProvider>
        <App />
      </MantineProvider>
    </React.StrictMode>
  );
} else {
  console.error('Root element with id "root" not found in the document');
}
