import { useState } from "react";

function useSub({ additionalData }) {
    const [status, setStatus] = useState('');
    const [message, setMessage] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        setStatus('loading');
        setMessage('');
        const finalFormEndpoint = "http://localhost:8080/";
        
        var id = e.target.name.value;
        var pointsValue = e.target.points.value;
        var points = 0
       
        var response = null
        var responseUpdate = null
        try {
            response = await fetch(finalFormEndpoint + "id/" + id, {
                headers: {
                    accept: 'application/json',
                },
            });

            const result = await response.json();

            if (!response.ok) {
                throw new Error(`Error! status: ${response.status}`);
            }
            console.log(result.points.Value)

            points = Number(result.points.Value) - Number(pointsValue)
            responseUpdate = await fetch(finalFormEndpoint + "points/" + points + "/id/"+ id, {
                method: 'Post',
                headers: {
                    accept: 'application/json',
                },
            });

            const resultUpdate = await responseUpdate.json();

            if (!responseUpdate.ok) {
                throw new Error(`Error! status: ${response.status}`);
            }


            setMessage(points);
            setStatus("updated");
        } catch (err) {
            setMessage(err.toString());
            setStatus('error');
        }
    };

    return { handleSubmit, status, message };
}

export default useSub;