import React, { useState } from "react";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faGlobeAmericas } from "@fortawesome/free-solid-svg-icons";

import { Tweets } from "./Tweets.js";

import { faEllipsisH } from "@fortawesome/free-solid-svg-icons";
import { faExchangeAlt } from "@fortawesome/free-solid-svg-icons";
import SideNav from "./SideNav.js";

const Tweet = () => {
  const [personsTweet, setTweets] = useState(Tweets);
  const [showNav, setNavShow] = useState(false);

  const toggle = () => {
    setNavShow(!showNav);
  };

  console.log(showNav);
  return (
    <>
      {showNav && <SideNav toggle={toggle} />}
      <section className="h-full lg:w-[580px] sm:w-full border-r-[1px] border-DarkGray lg:ml-[330px]">
        <div className="text-2xl w-full p-5 text-white font-bold flex justify-between items-center">
          <img
            src="https://www.gravatar.com/avatar/1b8fabaa8d66250a7049bdb9ecf44397?s=250&d=mm&r=x"
            alt="#"
            onClick={toggle}
            className="w-[45px] h-[45px] rounded-full hover:opacity-90 cursor-pointer lg:hidden sm:flex"
          />
          <h1 className="cursor-pointer sm:hidden lg:flex">Home</h1>
          <div className="text-white p-[8px]">
            <ion-icon name="logo-twitter" size="large"></ion-icon>
          </div>
          <div className="cursor-pointer hover:bg-DarkGray  hover:bg-opacity-20 rounded-full p-[10px]  flex items-center justify-center duration-150 ">
            <ion-icon name="sparkles-outline"></ion-icon>
          </div>
        </div>
        <article className="w-full lg:flex sm:hidden border-b-[1px] border-DarkGray">
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
              <FontAwesomeIcon icon={faGlobeAmericas}></FontAwesomeIcon>{" "}
              Everyone can reply
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
        {personsTweet.map((items) => {
          const { id, name, userName, image, likes, retweets, comment, tweet } =
            items;
          return (
            <article
              key={id}
              className="w-full flex p-3.5 hover:bg-DarkGray  hover:bg-opacity-10 duration-150 cursor-pointer border-b-[1px] border-DarkGray"
            >
              <div className="lg:w-1/6 sm:w-3/12 h-full flex items-center justify-center">
                <img
                  src={image}
                  alt="#"
                  className="mr-[5px] w-[50px] h-[50px] rounded-full hover:opacity-90 cursor-pointer"
                />
              </div>
              <div className="lg:w-5/6 sm:w-9/12">
                <header className="flex items-center w-full  p-1">
                  <h2 className=" md:flex sm:hidden text-white font-semibold lg:text-xl sm:text-lg hover:underline mr-[5px]">
                    {name}
                  </h2>
                  <h2 className=" sm:flex md:hidden text-white font-semibold lg:text-xl sm:text-lg hover:underline mr-[5px]">
                    {name.length > 10 ? `${name.substring(0, 9)}...` : name}{" "}
                  </h2>

                  <p className="text-DarkGray font-normal lg:text-lg sm:text-md no-underline mr-1">
                    {userName.length > 7
                      ? `${userName.substring(0, 6)}...`
                      : userName}{" "}
                    •
                  </p>

                  <p className="text-DarkGray font-normal hover:underline">
                    6m ago
                  </p>
                </header>
                <div className=" w-full p-1 mb-2">
                  <p className="text-white font-normal">{tweet}</p>
                </div>

                <div className="lg:w-11/12 sm:w-full flex items-center justify-between lg:p-2 sm:p-1">
                  <button className="flex items-center text-DarkGray hover:text-Primary">
                    <span className="text-DarkGray hover:text-Primary lg:text-2xl sm:text-xl hover:bg-Primary hover:bg-opacity-20 rounded-full p-[10px]  flex items-center justify-center duration-150 ">
                      <ion-icon name="chatbubble-outline"></ion-icon>
                    </span>
                    <p className="text-sm">{comment}</p>
                  </button>
                  <button className="flex items-center text-DarkGray hover:text-Green">
                    <span className="text-DarkGray hover:text-Green  lg:text-2xl sm:text-xl hover:bg-Green hover:bg-opacity-20 rounded-full px-[10px] py-[12px] flex items-center justify-center duration-150">
                      <FontAwesomeIcon icon={faExchangeAlt}></FontAwesomeIcon>
                    </span>
                    <p className="text-sm">{retweets}</p>
                  </button>
                  <button className="flex items-center text-DarkGray hover:text-Pink">
                    <span className="text-DarkGray hover:text-Pink lg:text-2xl sm:text-xl hover:bg-Pink hover:bg-opacity-20 rounded-full p-[10px] flex items-center justify-center duration-150 ">
                      <ion-icon name="heart-outline"></ion-icon>
                    </span>
                    <p className="text-sm ">{likes}</p>
                  </button>
                  <button>
                    <span>
                      <button>
                        <span className="text-DarkGray hover:text-Primary lg:text-2xl sm:text-xl hover:bg-Primary hover:bg-opacity-20 rounded-full p-[10px]  flex items-center justify-center duration-150 ">
                          <ion-icon name="share-social-outline"></ion-icon>
                        </span>
                      </button>
                    </span>
                  </button>
                </div>
              </div>
              <div className=" h-full text-DarkGray cursor-pointer hover:bg-DarkGray  hover:bg-opacity-20 rounded-full p-[10px]  flex items-center justify-center duration-150 ">
                <FontAwesomeIcon icon={faEllipsisH}></FontAwesomeIcon>
              </div>
            </article>
          );
        })}
      </section>
    </>
  );
};

export default Tweet;
