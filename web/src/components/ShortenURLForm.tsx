import { useState } from "react";
import { httpClient } from "../services/HttpClient";
import QRcode from "../assets/qrcode.svg";

export default function ShortenURLForm() {
  const [originalURL, setOriginalURL] = useState("");
  const [customAlias, setCustomAlias] = useState("");
  const [shortURL, setShortURL] = useState("");
  const [isCopied, setIsCopied] = useState(false);
  const [qrCodeBase64, setQrCodeBase64] = useState("");
  const [showModal, setShowModal] = useState(false);

  const toggleModal = async () => {
    await handleGenerateQRCode();
    setShowModal(!showModal);
  };

  const handleGenerateQRCode = async () => {
    const response = await httpClient.post("/qrcode", {
      url: originalURL,
    });
    if (response.status !== 200) {
      console.error("Failed to generate QR Code");
      return;
    }
    setQrCodeBase64("data:image/png;base64," + response.data.qr_code);
  };

  const copyToClipboard = async (url: string) => {
    console.log("copyToClipboard");
    try {
      await navigator.clipboard.writeText(url);
      setIsCopied(true);
    } catch (err) {
      console.error("Failed to copy text: ", err);
    }
  };

  const handleGenerateShortURL = async (event: any) => {
    event.preventDefault();
    const response = await httpClient.post("/shorten", {
      original_url: originalURL,
      custom_alias: customAlias,
    });

    if (response.status !== 200) {
      alert("invalid URL");
      console.error("Failed to shorten URL");
      return;
    }
    console.log(response);
    setShortURL(response.data.short_url);
    setIsCopied(false);
  };

  const downloadQRCode = () => {
    const byteCharacters = atob(qrCodeBase64.split(",")[1]);
    const byteNumbers = new Array(byteCharacters.length);
    for (let i = 0; i < byteCharacters.length; i++) {
      byteNumbers[i] = byteCharacters.charCodeAt(i);
    }
    const byteArray = new Uint8Array(byteNumbers);
    const blob = new Blob([byteArray], { type: "image/png" });
    const url = window.URL.createObjectURL(blob);
    const a = document.createElement("a");
    a.href = url;
    a.download = `${shortURL.split("/").pop()}.png`;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    window.URL.revokeObjectURL(url);
  };

  return (
    <div className="bg-gray-600 min-h-screen">
      <h1 className="pt-10 pb-4 flex justify-center text-5xl  text-white">
        Shorturl.life
      </h1>
      <h2 className="flex pb-6 justify-center text-sm text-gray-300">
        <p className="max-w-2xl text-center">
          Shorturl.life is a free URL shortener and QR code generator tool.
          Create your own short links and QR Codes in seconds. TEST
        </p>
      </h2>
      <div className="flex justify-center">
        <div className="bg-white shadow rounded-lg p-6 max-w-2xl w-full transition-transform hover:scale-105">
          <form onSubmit={handleGenerateShortURL} className="mb-4">
            <div className="mb-4">
              <label
                htmlFor="originalURL"
                className="block text-lg font-semibold mb-2"
              >
                Original URL
              </label>
              <input
                id="originalURL"
                type="url"
                value={originalURL}
                onChange={(e) => setOriginalURL(e.target.value)}
                required
                className="w-full px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring focus:border-blue-500"
              />
            </div>
            <div className="flex justify-between mb-4">
              <div className="w-5/12">
                <label
                  htmlFor="domain"
                  className="block text-lg font-semibold mb-2"
                >
                  Domain
                </label>
                <input
                  type="text"
                  readOnly
                  className="w-full px-4 py-2 rounded-lg border border-gray-300 bg-gray-200 focus:outline-none focus:ring focus:border-blue-500"
                  placeholder="http://localhost:5173"
                />
              </div>
              <p className="pt-10">/</p>
              <div className="w-5/12">
                <label
                  htmlFor="alias"
                  className="block text-lg font-semibold mb-2"
                >
                  Alias(optional)
                </label>
                <input
                  id="customAlias"
                  type="text"
                  value={customAlias}
                  onChange={(e) => setCustomAlias(e.target.value)}
                  className="w-full px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring focus:border-blue-500"
                />
              </div>
            </div>
            <div className="flex items-center justify-center pt-4">
              <button
                type="submit"
                className="py-2 px-6 bg-blue-500 hover:bg-blue-600 text-white font-semibold rounded-lg focus:outline-none focus:ring focus:ring-blue-300"
              >
                Shorten URL
              </button>
            </div>
            {shortURL && (
              <div>
                <div className="pt-6 flex  justify-between">
                  <div className="text-xl font-semibold flex pt-2">
                    <div>Shortened URL: </div>
                    <div className="pl-4 text-blue-500 underline items-center">
                      <a href={shortURL}>{shortURL}</a>
                    </div>
                  </div>
                </div>
                <div className="flex justify-center pt-4">
                  <div className="flex">
                    <img
                      src={QRcode}
                      alt="qrcode"
                      className="w-10 h-10 cursor-pointer"
                      onClick={toggleModal}
                    />
                  </div>
                  <div className="flex">
                    <button
                      type="button"
                      className="w-24 ml-2 text-xl bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-1 px-2 rounded focus:outline-none focus:shadow-outline"
                      onClick={() => copyToClipboard(shortURL)}
                    >
                      {isCopied ? "Copied!" : "Copy"}
                    </button>
                  </div>
                </div>
              </div>
            )}
          </form>
        </div>
      </div>
      {showModal && (
        <div className="fixed inset-0 bg-gray-700 bg-opacity-75 flex items-center justify-center">
          <div className="flex flex-col bg-white p-8 rounded-lg">
            <h2 className="text-2xl mb-4 flex justify-center">QR Code</h2>
            <div className="mb-4 w-50 h-50 flex justify-center">
              <img src={qrCodeBase64} alt="qrcode" />
            </div>
            <div className="flex justify-evenly">
              <button
                onClick={downloadQRCode}
                className="py-1 px-4 bg-gray-500 text-white rounded w-28"
              >
                Download
              </button>
              <button
                onClick={toggleModal}
                className="py-1 px-4 bg-gray-500 text-white rounded w-28"
              >
                Close
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
