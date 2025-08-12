'use client'

import Link from "next/link"

export default function Inputbox() {
    return(
        <div className="flex justify-center  py-20">
            <textarea className="bg-gray-900 w-[70%] h-96 rounded-xl p-4  text-white "
            value={textContent}
            onChange={(e) => setTextContent(e.target.value)}
            />
        </div>
    )
}