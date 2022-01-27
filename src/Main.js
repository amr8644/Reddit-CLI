import React from "react";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faGlobeAmericas } from "@fortawesome/free-solid-svg-icons";
import { faEllipsisH } from "@fortawesome/free-solid-svg-icons";
import { faRetweet } from "@fortawesome/free-solid-svg-icons";

const Tweet = () => {
  return (
    <section className="h-full w-[580px] border-r-[1px] border-DarkGray">
      <div className="text-2xl w-full p-5 sticky text-white font-bold flex justify-between items-center ">
        <h1 className="cursor-pointer">Home</h1>
        <div className="cursor-pointer hover:bg-DarkGray  hover:bg-opacity-20 rounded-full p-[10px]  flex items-center justify-center duration-150 ">
          <ion-icon name="sparkles-outline"></ion-icon>
        </div>
      </div>
      <article className="w-full flex border-b-[1px] border-DarkGray">
        <div className="w-1/6 h-full flex items-center justify-center">
          <img
            src="https://www.gravatar.com/avatar/1b8fabaa8d66250a7049bdb9ecf44397?s=250&d=mm&r=x"
            alt="#"
            className="mr-[5px] w-[70px] h-[70px] rounded-full hover:opacity-90 cursor-pointer"
          />
        </div>
        <div className="w-5/6 px-2 ">
          <textarea
            placeholder="What's happaning?"
            className="w-full h-[50px] bg-Body text-DarkGray focus:ring-0 outline-none border-none text-2xl "
          ></textarea>

          <p className="cursor-pointer w-[190px] flex justify-evenly items-center py-[3px] my-2 text-Primary hover:bg-Primary hover:bg-opacity-10 rounded-full font-semibold">
            <FontAwesomeIcon icon={faGlobeAmericas}></FontAwesomeIcon> Everyone
            can reply
          </p>
          <div className="w-full border-t-[1px] border-DarkGray flex items-center">
            <span className="cursor-pointer text-Primary hover:bg-Primary hover:bg-opacity-10 rounded-full my-3 w text-2xl w-[40px] h-[40px] flex items-center justify-center duration-150">
              <ion-icon name="image-outline"></ion-icon>
            </span>
            <span className="cursor-pointer text-Primary hover:bg-Primary hover:bg-opacity-10 rounded-full my-3 w text-2xl w-[40px] h-[40px] flex items-center justify-center duration-150">
              <ion-icon name="aperture-outline"></ion-icon>
            </span>
            <span className="cursor-pointer text-Primary hover:bg-Primary hover:bg-opacity-10 rounded-full my-3 w text-2xl w-[40px] h-[40px] flex items-center justify-center duration-150">
              <ion-icon name="bar-chart-outline"></ion-icon>
            </span>
            <span className="cursor-pointer text-Primary hover:bg-Primary hover:bg-opacity-10 rounded-full my-3 w text-2xl w-[40px] h-[40px] flex items-center justify-center duration-150">
              <ion-icon name="happy-outline"></ion-icon>
            </span>
            <span className="cursor-pointer text-Primary hover:bg-Primary hover:bg-opacity-10 rounded-full my-3 w text-2xl w-[40px] h-[40px] flex items-center justify-center duration-150">
              <ion-icon name="calendar-number-outline"></ion-icon>
            </span>
            <span className="cursor-pointer text-Primary hover:bg-Primary hover:bg-opacity-10 rounded-full my-3 w text-2xl w-[40px] h-[40px] flex items-center justify-center duration-150">
              <ion-icon name="location-outline"></ion-icon>
            </span>
            <button className="px-5 py-3 bg-Primary rounded-full hover:bg-opacity-70 text-white ml-24 font-semibold">
              Tweet
            </button>
          </div>
        </div>
      </article>
      <MainContent />
    </section>
  );
};

const MainContent = () => {
  return (
    <article className="w-full flex p-3.5 hover:bg-DarkGray  hover:bg-opacity-10 duration-150 cursor-pointer border-b-[1px] border-DarkGray">
      {/* User Picture */}
      <div className="w-1/6 h-full flex items-center justify-center">
        <img
          src="https://www.gravatar.com/avatar/1b8fabaa8d66250a7049bdb9ecf44397?s=250&d=mm&r=x"
          alt="#"
          className="mr-[5px] w-[50px] h-[50px] rounded-full hover:opacity-90 cursor-pointer"
        />
      </div>
      <div className="w-5/6">
        {/* Name and userName */}
        <header className="flex items-center w-full  p-1">
          <h2 className="text-white font-semibold text-xl hover:underline mr-[5px]">
            Premier League
          </h2>

          <p className="text-DarkGray font-normal text-lg no-underline mr-1">
            @premierleague •
          </p>

          <p className="text-DarkGray font-normal hover:underline"> 6m</p>
        </header>
        {/* Tweet Content */}
        <div className=" w-full p-1 mb-2">
          {" "}
          <p className="text-white font-normal">
            Last stadium you visited: <br /> Best stadium you've visited: <br />{" "}
            Dream stadium to visit:
          </p>
        </div>
        {/* Likes, Retweets , Comments and Share */}
        <div className="w-11/12 flex items-center justify-between p-2">
          <button>
            <span className="text-DarkGray hover:text-Primary text-2xl hover:bg-Primary hover:bg-opacity-20 rounded-full p-[10px]  flex items-center justify-center duration-150 ">
              <ion-icon name="chatbubble-outline"></ion-icon>
            </span>
          </button>
          <button>
            <span className="text-DarkGray hover:text-Green text-xl hover:bg-Green hover:bg-opacity-20 rounded-full px-[10px] py-[12px] flex items-center justify-center duration-150">
              <FontAwesomeIcon icon={faRetweet}></FontAwesomeIcon>
            </span>
          </button>
          <button>
            <span className="text-DarkGray hover:text-Pink text-2xl hover:bg-Pink hover:bg-opacity-20 rounded-full p-[10px] flex items-center justify-center duration-150 ">
              <ion-icon name="heart-outline"></ion-icon>
            </span>
          </button>
          <button>
            <span>
              <button>
                <span className="text-DarkGray hover:text-Primary text-2xl hover:bg-Primary hover:bg-opacity-20 rounded-full p-[10px]  flex items-center justify-center duration-150 ">
                  <ion-icon name="share-outline"></ion-icon>
                </span>
              </button>
            </span>
          </button>
        </div>
      </div>
      {/* More Button */}
      <div className=" h-full text-DarkGray cursor-pointer hover:bg-DarkGray  hover:bg-opacity-20 rounded-full p-[10px]  flex items-center justify-center duration-150 ">
        <FontAwesomeIcon icon={faEllipsisH}></FontAwesomeIcon>
      </div>
    </article>
  );
};

export default Tweet;
