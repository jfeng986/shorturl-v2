import { useState } from "react";
import { httpClient } from "../services/HttpClient";

export default function ShortenURLForm() {
  const [originalURL, setOriginalURL] = useState("");
  const [customAlias, setCustomAlias] = useState("");
  const [shortURL, setShortURL] = useState("");
  const [isCopied, setIsCopied] = useState(false);

  const copyToClipboard = async (url: string) => {
    try {
      await navigator.clipboard.writeText(url);
      setIsCopied(true);
    } catch (err) {
      console.error("Failed to copy text: ", err);
    }
  };

  const handleSubmit = async (event: any) => {
    event.preventDefault();
    const response = await httpClient.post("/shorten", {
      original_url: originalURL,
      custom_alias: customAlias,
    });

    if (response.status !== 200) {
      console.error("Failed to shorten URL");
      return;
    }
    console.log(response);
    setShortURL(response.data.short_url);
    setIsCopied(false);
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100 text-gray-800">
      <h1 className="text-6xl font-bold mb-5">GoShort.life</h1>
      <div className="bg-white shadow rounded-lg px-8 py-6 max-w-2xl w-full transition-all ease-in-out duration-500 transform hover:scale-105">
        <form onSubmit={handleSubmit}>
          <div className="mb-5">
            <label
              className="block text-xl font-semibold mb-2"
              htmlFor="originalURL"
            >
              Original URL
            </label>
            <input
              className="shadow appearance-none border border-gray-300 rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="originalURL"
              type="url"
              value={originalURL}
              onChange={(e) => setOriginalURL(e.target.value)}
              required
            />
          </div>
          <div className="mb-5">
            <label
              className="block text-xl font-semibold mb-2"
              htmlFor="customAlias"
            >
              Custom Alias (optional)
            </label>
            <input
              className="shadow appearance-none border border-gray-300 rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="customAlias"
              type="text"
              value={customAlias}
              onChange={(e) => setCustomAlias(e.target.value)}
            />
          </div>
          <div className="flex items-center justify-center p-5">
            <button
              className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
              type="submit"
            >
              Shorten URL
            </button>
          </div>
        </form>
        {shortURL && (
          <div className="mt-4 flex  justify-between">
            <div className="text-xl font-semibold flex">
              <div>Shortened URL: </div>
              <div className="pl-4 text-blue-500 underline items-center">
                <a href={shortURL} target="_blank" rel="noreferrer">
                  {shortURL}
                </a>
              </div>
            </div>

            <div>
              <button
                className="ml-2 text-xl bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-1 px-2 rounded focus:outline-none focus:shadow-outline"
                onClick={() => copyToClipboard(shortURL)}
              >
                {isCopied ? "Copied!" : "Copy"}
              </button>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
