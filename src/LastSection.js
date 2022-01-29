import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

import { faSearch } from "@fortawesome/free-solid-svg-icons";

const LastSection = () => {
  return (
    <section className="bg-Body text-white h-screen  lg:w-[450px] sm:hidden lg:flex items-center  flex-col border-r-[1px] border-DarkGray lg:fixed lg:ml-[910px]">
      <div className=" relative w-full flex items-center justify-center p-3 ">
        <FontAwesomeIcon
          icon={faSearch}
          className="absolute left-[40px] text-DarkGray text-xl"
        ></FontAwesomeIcon>
        <input
          type="search"
          placeholder="Search Twitter..."
          className="text-md rounded-full outline-none w-11/12 h-[45px] pl-10 bg-DarkGray bg-opacity-20 text-LightGray focus:ring-2 focus:ring-Primary focus:bg-Body focus:bg-opacity-100"
        />
      </div>
      <article className="bg-DarkGray bg-opacity-20 w-11/12 mt-5 rounded-xl">
        <h2 className="text-2xl font-bold py-3 px-2">Who to follow</h2>
        <div className="cursor-pointer w-full py-3 my-4 hover:bg-DarkGray hover:bg-opacity-20  duration-150 flex items-center  justify-evenly">
          <img
            src="https://images.beinsports.com/o0R8LGIyrAy1lqFErZd-RMGcADo=/full-fit-in/1000x0/jamie-vardy-cropped_1lo75eler0py015yapewdietqr.jpg"
            alt="#"
            className="mr-[5px] w-[55px] h-[55px] rounded-full"
          />
          <div className="mx-[5px]">
            <h4 className="font-bold">Vardy</h4>
            <p className="text-DarkGray">@vardy09</p>
          </div>
          <button className="px-8 font-semibold py-2 bg-white text-Body  rounded-full hover:bg-opacity-70">
            Follow
          </button>
        </div>
        <div className="cursor-pointer w-full py-3 my-4 hover:bg-DarkGray hover:bg-opacity-20  duration-150 flex items-center  justify-evenly">
          <img
            src="https://encrypted-tbn3.gstatic.com/images?q=tbn:ANd9GcR_BSXPlBjoBeJruSaCamv7kQuMNjoIIWX0CITXUVoapFCbRM9g"
            alt="#"
            className="mr-[5px] w-[55px] h-[55px] rounded-full"
          />
          <div className="mx-[5px]">
            <h4 className="font-bold">Messi</h4>
            <p className="text-DarkGray">@Messi10</p>
          </div>
          <button className="px-8 font-semibold py-2 bg-white text-Body  rounded-full hover:bg-opacity-70">
            Follow
          </button>
        </div>
      </article>
    </section>
  );
};

export default LastSection;
