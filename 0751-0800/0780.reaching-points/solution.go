package reachingpoints

// func reachingPoints(sx int, sy int, tx int, ty int) bool {
// 	for {
// 		if tx < sx || ty < sy {
// 			return false
// 		}
// 		if tx == sx && ty == sy {
// 			return true
// 		}

// 		if tx == ty {
// 			return false
// 		}

// 		for tx > ty && tx >= sx {
// 			tx -= ty * ((tx - sx) / ty)
// 			// tx >= sx
// 			if tx == sx && ty == sy {
// 				return true
// 			}
// 		}
// 		for ty > tx && ty >= sy {
// 			ty -= tx * ((ty - sy) / tx)
// 			// ty >= sy
// 			if tx == sx && ty == sy {
// 				return true
// 			}
// 		}
// 	}
// }

func reachingPoints(sx, sy, tx, ty int) bool {
	for tx > sx && ty > sy && tx != ty {
		if tx > ty {
			tx %= ty
		} else {
			ty %= tx
		}
	}
	switch {
	case tx == sx && ty == sy:
		return true
	case tx == sx:
		return ty > sy && (ty-sy)%tx == 0
	case ty == sy:
		return tx > sx && (tx-sx)%ty == 0
	default:
		return false
	}
}
