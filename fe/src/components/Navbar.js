import React from 'react';
import { Link } from 'react-router-dom';

const Navbar = () => {
  const styles = {
    nav: { 
      display: 'flex', 
      justifyContent: 'space-between', 
      alignItems: 'center',
      padding: '10px 20px', 
      background: '#333', 
      color: '#fff' 
    },
    logo: { 
      fontWeight: 'bold',
      marginRight: 'auto'
    },
    link: { 
      margin: '0 10px', 
      color: '#61dafb', 
      textDecoration: 'none' 
    },
    linkContainer: {
      display: 'flex',
      alignItems: 'center'
    }
  };

  return (
    <nav style={styles.nav}>
      <a href="/" style={styles.logo}>
        <h1>LiveScore</h1>
      </a>
      <div style={styles.linkContainer}>
        <Link to="/login" style={styles.link}>Login</Link>
        <Link to="/signup" style={styles.link}>Sign Up</Link>
      </div>
    </nav>
  );
};

export default Navbar;