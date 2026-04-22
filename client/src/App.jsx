import { use, useEffect, useState } from 'react'
import './App.css'

function App() {
  const [name, setName] = useState("");
  const [users, setUsers] = useState([]);

  async function loadUsers() {
    const response = await fetch(import.meta.env.VITE_API + '/users')
    const data = await response.json();
    setUsers(data.users);
  }

  useEffect(() => {
    loadUsers();
  }, [])


  const handleSubmit = async (e) => {
    e.preventDefault();

    const response = await fetch(import.meta.env.VITE_API + '/users', {
      method: 'POST',
      body: JSON.stringify({ name }),
      headers: {
        "Content-Type": "application/json"
      }
    });

    const data = await response.json();
    loadUsers();

  }



  return (
    <div>
      <form onSubmit={handleSubmit} >
        <input type="text" placeholder='Coloca tu nombre'
          onChange={(e) => setName(e.target.value)}
        />
        <button>Guardar</button>
      </form>

      <ul>
        {users.map(user =>
          <li key={user.name}>{user.name}</li>

        )}
      </ul>
    </div>
  )
}

export default App
