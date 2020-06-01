import React from 'react'
import '../style/styles.css'
import avatars from '../img/avatars.svg'

class Main extends React.Component {

  constructor() {
    super()
    this.state = {
      loading: false,
      randomFact: null
    }
  }

  async componentDidMount() {

    const url = "/api/v1/trivia"
    const response = await fetch(url, {
      method: 'get',
      headers: {
        "Content-Type": "application/json"
      }
    })

    const data = await response.json()

    this.setState({randomFact: data.trivia})

  }

  render() {

    const hours = new Date().getHours()
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
            
            <p className="masthead-subheading font-weight-light mb-0">Wanna know a random fact that ever happened on this date?</p>
            {this.state.loading || !this.state.randomFact ? 
              <p className="masthead-subheading font-weight-light mb-0">fetching a random fact</p>
              :
              <p className="masthead-subheading font-weight-light mb-0">{this.state.randomFact}</p>
            }
        </div>
      </main>
    )

  }
}

export default Main