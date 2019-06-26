/** @jsx preact.h */

// A simple click counter component, can be a functional compontent as well
class Counter extends preact.Component {
	constructor(props) {
    super(props)
    color.new()
  }
  
  doColorChange() {
    color.changec();
    render();
  }

  render() {
    let styles = {
      color: color.data.value
    }

    return (
      <div>
        <div style={styles}>count:{counter.data.value} and color {color.data.value}</div>
				<button onClick={() => counter.add(1)}>Plus 1</button>
        <button onClick={() => counter.substract(1)}>Moins 1</button>
        <button onClick={() => { this.doColorChange()}}>Change Color</button>
      </div>
    )
  }
}

// Render top-level component, pass controller data as props
const render = () => {
  preact.render(<Counter />, document.getElementById('app'), document.getElementById('app').lastElementChild);
}
	
// Call global render() when controller changes
counter.render = render;
// Render of the first time
render();
