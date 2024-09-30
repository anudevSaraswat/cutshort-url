import React, { useState } from "react";
import axios from 'axios';

const CutshortURL = () => {
    const [longUrl, setLongUrl] = useState("");
    const [shortUrl, setShortUrl] = useState("");

    const apiURL = process.env.REACT_APP_API_URL;

    // Mock URL shortening function (you can replace this with an actual API call)
    const shortenUrl = (url) => {
        const urlReq = {
            url: url,
        }
        axios.post(`${apiURL}/api/shorten`, urlReq).then((resp) => setShortUrl(resp.data.url))
            .catch((error) => alert(error));
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        if (longUrl) {
            const shortened = shortenUrl(longUrl);
            setShortUrl(shortened);
        }
    };

    return (
        <div className="url-shortener">
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    placeholder="Enter a long URL"
                    value={longUrl}
                    onChange={(e) => setLongUrl(e.target.value)}
                    required
                />
                <button type="submit">Shorten URL</button>
            </form>
            {shortUrl && (
                <div className="result">
                    <p>Shortened URL:</p>
                    <a href={`${apiURL}/api/resolve/${shortUrl}`} target="_blank" rel="noopener noreferrer">
                        {shortUrl}
                    </a>
                </div>
            )}
        </div>
    );
};

export default CutshortURL;
