package _5_select_data_flow

func MultiplexingExample(stopCh chan struct{},
	inputChA, inputChB chan int,
	outputChC, outputChD chan int) {

	for {
		var data int
		var isOpen bool

		// read from whichever channel has data
		select {
		case data, isOpen = <-inputChA:
			if !isOpen {
				inputChA = nil

				if inputChB == nil {
					// both channels are closed
					return
				}
			}

		case data, isOpen = <-inputChB:
			if !isOpen {
				// disable this case
				inputChB = nil

				if inputChA == nil {
					// both channels are closed
					return
				}
			}

		case <-stopCh:
			// give up
			return
		}

		// write to whichever channel is empty
		select {
		case outputChC <- data:

		case outputChD <- data:

		case <-stopCh:
			// give up
			return
		}
	}
}
