package main

const (
	INLINE_VIDEO = 1 << iota
	INLINE_TEXT
	INLINE_CHAT
)

func main() {

	println(INLINE_VIDEO)
	println(INLINE_TEXT)

	inlineHelps := 0
	inlineHelps |= INLINE_VIDEO
	inlineHelps |= INLINE_TEXT
	inlineHelps |= INLINE_CHAT

	println(inlineHelps)

	inlineHelps ^= INLINE_CHAT

	hasVideo := inlineHelps & INLINE_VIDEO

	println(inlineHelps)
	println(hasVideo)
}
