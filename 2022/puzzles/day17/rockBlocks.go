package day17

// Rock blocks define different boundaries for moving in certain directions, these
// are the parts of the rock that could hit something when moving in that direction.
type RockBlock interface {
	leftBoundary() []point
	rightBoundary() []point
	downBoundary() []point
	// allCoords must return points making up block, ordered from lowest to highest
	allCoords() []point

	moveLeft() RockBlock
	moveRight() RockBlock
	moveDown() RockBlock
	moveUp() RockBlock
}

/*
Rock block 1 - Horizontal Line: ####
*/
type Rock1 struct {
	position point
}

func (r Rock1) leftBoundary() []point {
	return []point{r.position}
}

func (r Rock1) rightBoundary() []point {
	return []point{point{r.position.x + 3, r.position.y}}
}

func (r Rock1) downBoundary() []point {
	boundary := make([]point, 0)
	for i := r.position.x; i < r.position.x+4; i++ {
		boundary = append(boundary, point{i, r.position.y})
	}
	return boundary
}

func (r Rock1) allCoords() []point {
	return r.downBoundary()
}

// TODO how to share this across multiple types?
func (r Rock1) moveLeft() RockBlock {
	r.position.x--
	return r
}
func (r Rock1) moveRight() RockBlock {
	r.position.x++
	return r
}
func (r Rock1) moveDown() RockBlock {
	r.position.y--
	return r
}
func (r Rock1) moveUp() RockBlock {
	r.position.y++
	return r
}

// TODO wtf is going on with gofmt and block comments??
/*
Rock block 2 - Cross:

	 #
	###
	 #
*/
type Rock2 struct {
	position point
}

func (r Rock2) leftBoundary() []point {
	// left point
	boundary := []point{point{r.position.x, r.position.y + 1}}
	// vertical middle bit
	for j := r.position.y; j < r.position.y+3; j++ {
		boundary = append(boundary, point{r.position.x + 1, j})
	}
	return boundary
}

func (r Rock2) rightBoundary() []point {
	// right point
	boundary := []point{point{r.position.x + 2, r.position.y + 1}}
	// vertical middle bit
	for j := r.position.y; j < r.position.y+3; j++ {
		boundary = append(boundary, point{r.position.x + 1, j})
	}
	return boundary
}

func (r Rock2) downBoundary() []point {
	// bottom point
	boundary := []point{point{r.position.x + 1, r.position.y}}
	// horizontal middle bit
	for i := r.position.x; i < r.position.x+3; i++ {
		boundary = append(boundary, point{i, r.position.y + 1})
	}
	return boundary
}

func (r Rock2) allCoords() []point {
	return append(r.downBoundary(), point{r.position.x + 1, r.position.y + 2})
}

func (r Rock2) moveLeft() RockBlock {
	r.position.x--
	return r
}
func (r Rock2) moveRight() RockBlock {
	r.position.x++
	return r
}
func (r Rock2) moveDown() RockBlock {
	r.position.y--
	return r
}
func (r Rock2) moveUp() RockBlock {
	r.position.y++
	return r
}

/*
Rock block 3 - Backwards L:

	  #
	  #
	###
*/
type Rock3 struct {
	position point
}

func (r Rock3) leftBoundary() []point {
	return []point{
		r.position,
		point{r.position.x + 2, r.position.y + 1},
		point{r.position.x + 2, r.position.y + 2},
	}
}

func (r Rock3) rightBoundary() []point {
	boundary := make([]point, 0)
	for j := r.position.y; j < r.position.y+3; j++ {
		boundary = append(boundary, point{r.position.x + 2, j})
	}
	return boundary
}

func (r Rock3) downBoundary() []point {
	boundary := make([]point, 0)
	for i := r.position.x; i < r.position.x+3; i++ {
		boundary = append(boundary, point{i, r.position.y})
	}
	return boundary
}

func (r Rock3) allCoords() []point {
	return append(
		r.downBoundary(),
		point{r.position.x + 2, r.position.y + 1},
		point{r.position.x + 2, r.position.y + 2},
	)
}

func (r Rock3) moveLeft() RockBlock {
	r.position.x--
	return r
}
func (r Rock3) moveRight() RockBlock {
	r.position.x++
	return r
}
func (r Rock3) moveDown() RockBlock {
	r.position.y--
	return r
}
func (r Rock3) moveUp() RockBlock {
	r.position.y++
	return r
}

/*
	 Rock block 4 - Vertical Line: #
								 						     #
		                             #
		                             #
*/
type Rock4 struct {
	position point
}

func (r Rock4) leftBoundary() []point {
	boundary := make([]point, 0)
	for j := r.position.y; j < r.position.y+4; j++ {
		boundary = append(boundary, point{r.position.x, j})
	}
	return boundary
}

func (r Rock4) rightBoundary() []point {
	return r.leftBoundary()
}

func (r Rock4) downBoundary() []point {
	return []point{r.position}
}

func (r Rock4) allCoords() []point {
	return r.leftBoundary()
}

func (r Rock4) moveLeft() RockBlock {
	r.position.x--
	return r
}
func (r Rock4) moveRight() RockBlock {
	r.position.x++
	return r
}
func (r Rock4) moveDown() RockBlock {
	r.position.y--
	return r
}
func (r Rock4) moveUp() RockBlock {
	r.position.y++
	return r
}

/*
	 Rock block 5 - Square: ##
						  					  ##
*/
type Rock5 struct {
	position point
}

func (r Rock5) leftBoundary() []point {
	return []point{r.position, point{r.position.x, r.position.y + 1}}
}

func (r Rock5) rightBoundary() []point {
	return []point{
		point{r.position.x + 1, r.position.y},
		point{r.position.x + 1, r.position.y + 1},
	}
}

func (r Rock5) downBoundary() []point {
	return []point{r.position, point{r.position.x + 1, r.position.y}}
}

func (r Rock5) allCoords() []point {
	return append(r.leftBoundary(), r.rightBoundary()...)
}

func (r Rock5) moveLeft() RockBlock {
	r.position.x--
	return r
}
func (r Rock5) moveRight() RockBlock {
	r.position.x++
	return r
}
func (r Rock5) moveDown() RockBlock {
	r.position.y--
	return r
}
func (r Rock5) moveUp() RockBlock {
	r.position.y++
	return r
}

type point struct {
	x int
	y int
}
