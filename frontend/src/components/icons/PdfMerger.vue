<template>
  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none">
    <!-- First PDF (back) -->
    <path d="M8 6a1 1 0 011-1h4l4 4v8a1 1 0 01-1 1H9a1 1 0 01-1-1V6z" stroke="currentColor" stroke-width="1.5" fill="none" opacity="0.5"/>
    
    <!-- Second PDF (front) -->
    <path d="M4 8a1 1 0 011-1h4l4 4v8a1 1 0 01-1 1H5a1 1 0 01-1-1V8z" stroke="currentColor" stroke-width="1.5" fill="none"/>
    <path d="M9 7v4h4" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
    
    <!-- Merge arrows -->
    <path d="M16 10l3 3-3 3" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
    <path d="M16 16l3-3" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
    
    <!-- Result PDF -->
    <rect x="19" y="9" width="3" height="6" rx="0.5" stroke="currentColor" stroke-width="1.5" fill="none"/>
    <path d="M20 11h1M20 13h1" stroke="currentColor" stroke-width="1" stroke-linecap="round"/>
  </svg>
</template>
