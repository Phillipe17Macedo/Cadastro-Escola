import React, { useState, useEffect } from 'react';
import axios from 'axios';

function Professores() {
  const [professores, setProfessores] = useState([]);
  const [novoProfessor, setNovoProfessor] = useState({ nome: '', email: '', cpf: '' });

  useEffect(() => {
    axios.get('http://localhost:8080/professores')
      .then(response => setProfessores(response.data))
      .catch(error => console.log(error));
  }, []);

  const handleChange = (e) => {
    setNovoProfessor({ ...novoProfessor, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    axios.post('http://localhost:8080/professores', novoProfessor)
      .then(response => {
        setProfessores([...professores, response.data]);
        setNovoProfessor({ nome: '', email: '', cpf: '' });
      })
      .catch(error => console.log(error));
  };

  return (
    <div>
      <h2>Professores</h2>
      <ul>
        {professores.map(professor => (
          <li key={professor.ID}>{professor.Nome} - {professor.Email}</li>
        ))}
      </ul>

      <h3>Cadastrar Novo Professor</h3>
      <form onSubmit={handleSubmit}>
        <input type="text" name="nome" placeholder="Nome" value={novoProfessor.nome} onChange={handleChange} required />
        <input type="email" name="email" placeholder="Email" value={novoProfessor.email} onChange={handleChange} required />
        <input type="text" name="cpf" placeholder="CPF" value={novoProfessor.cpf} onChange={handleChange} required />
        <button type="submit">Cadastrar</button>
      </form>
    </div>
  );
}

export default Professores;