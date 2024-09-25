import React, { useState } from "react";

const CutshortURL = () => {
    const [longUrl, setLongUrl] = useState("");
    const [shortUrl, setShortUrl] = useState("");

    // Mock URL shortening function (you can replace this with an actual API call)
    const shortenUrl = (url) => {
        return `https://short.ly/${Math.random().toString(36).substr(2, 5)}`;
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
                    <a href={shortUrl} target="_blank" rel="noopener noreferrer">
                        {shortUrl}
                    </a>
                </div>
            )}
        </div>
    );
};

export default CutshortURL;
