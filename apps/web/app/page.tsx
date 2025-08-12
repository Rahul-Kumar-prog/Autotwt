'use client'
import Image from "next/image";
import { Card } from "@repo/ui/card";
import { Gradient } from "@repo/ui/gradient";
import { TurborepoLogo } from "@repo/ui/turborepo-logo";
import { useState } from "react";
import { Content } from "next/font/google";
import { text } from "stream/consumers";
import Navbar from "./component/layout/navbar";
import Inputbox from "./component/ui/inputbox";
import Slidebar from "./component/layout/slidebar";


export default function Page() {
  const [textContent, setTextContent] = useState('');
  const [responseMessage, setResponseMessage] = useState('');
  // const [isLoading, setIsLoading] = useState(false);

  const handlePost = async () => {
    if (!textContent.trim()) {
      alert('please enter sometext before posting!')
      return;
    }
    // setIsLoading(true);

    try {
      console.log('Attempting to connect to server....')
      console.log('payload:', { content: textContent })

      const apiUrl = 'http://localhost:8000';
      const response = await fetch(`${apiUrl}/api/post`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          Content: textContent
        })
      });

      console.log('Response status:', response.status);
      console.log('Response headers:', response.headers);

      if (response.ok) {
        const data = await response.json();
        console.log('Response data:', data)
        setResponseMessage(data.message);
        setTextContent('');
      } else {
        const errorText = await response.text();
        console.log('error response:', errorText)
        setResponseMessage('Error : failed to post')
      }
    } catch (error) {
      console.error('Error posting:', error);
      setResponseMessage('error could not connect to server')
    } finally {
      // setIsLoading(false);
    }
  }
  return (
    <main className="">
      <Navbar/>
      {/* <Inputbox/> */}
      {/* <Slidebar/> */}

      <div className="flex justify-center  py-20 ">
            <textarea className="bg-gray-900 w-[70%] h-96 rounded-xl p-4  text-white "
            value={textContent}
            onChange={(e) => setTextContent(e.target.value)}
            />
        </div>

      <div className="flex justify-center items-center">
        <button
          className={`bg-blue-600 w-[20%] font-bold h-10 rounded-full text-white cursor-pointer hover:bg-blue-700 transition-colors          }`}
          onClick={handlePost}
          // disabled={isLoading}
        >
          Post
          {/* {isLoading ? 'Posting...' : 'Post'} */}
        </button>
      </div>


      {/* Display response message */}
      {responseMessage && (
        <div className="flex justify-center mt-6">
          <div className={`p-4 rounded-lg ${responseMessage.includes('Error')
              ? 'bg-red-100 text-red-700 border border-red-300'
              : 'bg-green-100 text-green-700 border border-green-300'
            }`}>
            {responseMessage}
          </div>
        </div>
      )}
    </main>
  );
}
