import { useState } from "react";

function useForm({ additionalData }) {
    const [status, setStatus] = useState('');
    const [message, setMessage] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        setStatus('loading');
        setMessage('');

        const finalFormEndpoint = "http://localhost:8080/";
        const data = Array.from(e.target.elements)
            .filter((input) => input.name)
            .reduce((obj, input) => Object.assign(obj, { [input.name]: input.value }), {});

        if (additionalData) {
            Object.assign(data, additionalData);
        }
        var response = null
        console.log(data.name)
        try {
            response = await fetch(finalFormEndpoint + "id/" + data.name, {
                method: 'GET',
                headers: {
                    accept: 'application/json',
                },
            });

            const result = await response.json();

            if (!response.ok) {
                throw new Error(`Error! status: ${response.status}`);
            }

            setMessage(result);
            setStatus("success");
        } catch (err) {
            setMessage(err.toString());
            setStatus('error');
        }
    };

    return { handleSubmit, status, message };
}

export default useForm;