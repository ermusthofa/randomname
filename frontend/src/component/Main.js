import React from 'react'
import '../style/styles.css'
import avatars from '../img/avatars.svg'

function Main() {

  const date = new Date()
  const hours = date.getHours()
  let timeOfDay

  if (hours < 12) {
    timeOfDay = "Morning"
  } else if (hours >= 12 && hours < 18) {
    timeOfDay = "Afternoon"
  } else {
    timeOfDay = "Night"
  }

  return (
    <main className="masthead bg-primary text-white text-center">
      <div className="container d-flex align-items-center flex-column">
          <img className="masthead-avatar mb-5" src={avatars} alt="" />
          <h1 className="masthead-heading mb-0">Hi! Good {timeOfDay}!</h1>
          
          <div className="divider-custom divider-light">
              <div className="divider-custom-line"></div>
              <div className="divider-custom-icon"><i className="fas fa-star"></i></div>
              <div className="divider-custom-line"></div>
          </div>
          
          <p className="masthead-subheading font-weight-light mb-0">Wanna know what ever happened in today's date?</p>
          <p className="masthead-subheading font-weight-light mb-0">Today, is blablabla</p>
      </div>
    </main>
  )
}

export default Main