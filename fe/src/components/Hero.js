import React from 'react';

const Hero = () => {
  const styles = {
    hero: { textAlign: 'center', padding: '50px 20px', background: '#f0f0f0' }
  };
  return (
    <div style={styles.hero}>
      <h1>Welcome to LiveScore</h1>
      <p>Live score platform to update your favorite match!</p>
    </div>
  );
};


export default Hero;