'use client'

import Link from "next/link"
import Slidebar from "./slidebar"



export default function Navbar() {
    const Popup = () => {
        console.log("popup:open")
    }
    return (
        <div className="w-full h-18 bg-gray-900 items-center justify-between flex">
            <div className="text-4xl ml-8 cursor-pointer">Autotwt</div>

            <button onClick={Popup} className="w-8 h-8 bg-white mr-8 items-center flex flex-col justify-center place-content-between gap-[3] rounded-tl-md rounded-br-md cursor-pointer  hover:scale-120 duration-300 transition-all">
              <div className="w-5 h-1 rounded-2xl bg-gray-900"></div>
              <div className="w-5 h-1 rounded-2xl bg-gray-900"></div>
              <div className="w-5 h-1 rounded-2xl bg-gray-900"></div>
            </button>
          </div>
    )
}
