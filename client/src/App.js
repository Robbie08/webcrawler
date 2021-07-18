import React, { useState, useCallback } from 'react';
import styled from 'styled-components';
import Input from '@material-ui/core/Input';
import Button from '@material-ui/core/Button';

const Container = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
`;

const EntryContainer = styled.div`
  display: flex;
  flex-direction: row;
`;

const StyledInput = styled(Input)`
  margin-right: 10px;
`;

function App() {
  const [input, setInput] = useState('');
  
  const handleOnClick = useCallback(() => {
      alert(input);
  }, [input]);

  return (
    <Container>
      <p>test</p>
      <EntryContainer>
        <StyledInput
          color="secondary"
          onChange={(e) => setInput(e.target.value)}
          
        />
        <Button variant="contained" color="primary" onClick={handleOnClick} >Enter</Button>
      </EntryContainer>
    </Container>
  );
}

export default App;
