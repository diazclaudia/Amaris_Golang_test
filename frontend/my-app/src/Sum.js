import useSum from "./UseSum";
import { Button } from 'react-bootstrap';

import Form from 'react-bootstrap/Form';

const Sum = () => {
  const additionalData = {
    sent: new Date().toISOString(),
  };

  const { handleSubmit, status, message } = useSum({
    additionalData,
  });

  if (status === "updated") {
    return (
      <div style={{backgroundColor: message===500? "yellow":"none"}}>
      <h1 style={{display: message===500? "":"none"}}>Has ganado un premio</h1>
      <h1><strong>Detalle de puntos</strong></h1><br/>
        <p>POINTS: {message}</p>
      </div>
    );
  }

  if (status === "error") {
    return (
      <>
        <div className="text-2xl">Something bad happened!</div>
        <div className="text-md">{message}</div>
      </>
    );
  }

  return (
    <Form
      onSubmit={handleSubmit}
      method="POST"
    >
      <div className="pt-0 mb-3">
        <input
          type="text"
          placeholder="Your ID"
          name="name"
          className="focus:outline-none focus:ring relative w-full px-3 py-3 text-sm text-gray-600 placeholder-gray-400 bg-white border-0 rounded shadow outline-none"
        />
        <input
          type="number"
          placeholder="Puntos a sumar"
          name="points"
          className="focus:outline-none focus:ring relative w-full px-3 py-3 text-sm text-gray-600 placeholder-gray-400 bg-white border-0 rounded shadow outline-none"
        />
      </div>
      {status !== "loading" && (
        <div className="pt-0 mb-3">
          <Button
            className="active:bg-blue-600 hover:shadow-lg focus:outline-none px-6 py-3 mb-1 mr-1 text-sm font-bold text-white uppercase transition-all duration-150 ease-linear bg-blue-500 rounded shadow outline-none"
            type="submit"
          >
            Sumar puntos
          </Button>
        </div>
      )}
    </Form>
    
  );
};

export default Sum;