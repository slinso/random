package uuid_test

/*
CPU Intel(R) Core(TM) i3-4160T CPU @ 3.10GHz

BenchmarkShortUUID-4                      118107             10364 ns/op            2953 B/op        136 allocs/op
BenchmarkUUID-4                           794074              1444 ns/op              16 B/op          1 allocs/op
BenchmarkUUIDString-4                     741462              1623 ns/op              64 B/op          2 allocs/op
BenchmarkXid-4                          15945531                72.7 ns/op             0 B/op          0 allocs/op
BenchmarkXidString-4                    11560804               102 ns/op               0 B/op          0 allocs/op
BenchmarkKsuid-4                          789439              1501 ns/op               0 B/op          0 allocs/op
BenchmarkKsuidString-4                    639405              1824 ns/op               0 B/op          0 allocs/op
BenchmarkBetterGUID-4                    8451936               140 ns/op              32 B/op          1 allocs/op
BenchmarkUlidFixedEntropy-4             18076548                61.3 ns/op            16 B/op          1 allocs/op
BenchmarkUlidFixedEntropyString-4       11225040               103 ns/op              16 B/op          1 allocs/op
BenchmarkUlidRandomEverytime-4           5940288               187 ns/op              64 B/op          2 allocs/op
BenchmarkSonyflake-4                       31450             38760 ns/op               2 B/op          0 allocs/op
BenchmarkSid-4                           3341583               354 ns/op             115 B/op          3 allocs/op
BenchmarkUUIDv4RFC4122String-4            776656              1583 ns/op              64 B/op          2 allocs/op
BenchmarkUUIDv4Raw-4                      748167              1446 ns/op              16 B/op          1 allocs/op
BenchmarkGONanoID-4                       564336              1855 ns/op             160 B/op          3 allocs/op
BenchmarkGONanoIDCustom-4                 147259              7528 ns/op            1168 B/op          4 allocs/op
*/

import (
	crypto_rand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"

	"testing"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/chilts/sid"
	"github.com/gofrs/uuid"
	guuid "github.com/google/uuid"
	"github.com/kjk/betterguid"
	"github.com/lithammer/shortuuid"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/oklog/ulid"
	"github.com/rs/xid"
	"github.com/segmentio/ksuid"
	"github.com/sony/sonyflake"
)

// BenchmarkShortUUID "github.com/lithammer/shortuuid"
func BenchmarkShortUUID(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = shortuuid.New()
	}
}

// BenchmarkUUID "github.com/google/uuid"
func BenchmarkUUID(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = guuid.New()
	}
}

// BenchmarkUUIDString "github.com/google/uuid"
func BenchmarkUUIDString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = guuid.New().String()
	}
}

// BenchmarkXid "github.com/rs/xid"
func BenchmarkXid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = xid.New()
	}
}

func TestXid(t *testing.T) {
	for i := 0; i < 10; i++ {
		id := xid.New()
		fmt.Println("Xid:", id)
	}
}

// BenchmarkXidString "github.com/rs/xid"
func BenchmarkXidString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = xid.New().String()
	}
}

// BenchmarkKsuid "github.com/segmentio/ksuid"
func BenchmarkKsuid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = ksuid.New()
	}
}

// BenchmarkKsuidString "github.com/segmentio/ksuid"
func BenchmarkKsuidString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = ksuid.New().String()
	}
}

// BenchmarkBetterGUID "github.com/kjk/betterguid"
func BenchmarkBetterGUID(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = betterguid.New()
	}
}

func TestBetterGUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		id := betterguid.New()
		fmt.Println("BetterGUID:", id)
	}
}

// BenchmarkUlidFixedEntropy "github.com/oklog/ulid" with fixed entropy
func BenchmarkUlidFixedEntropy(b *testing.B) {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = ulid.MustNew(ulid.Timestamp(t), entropy)
	}
}

func TestUlid(_ *testing.T) {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))

	for i := 0; i < 10; i++ {
		id := ulid.MustNew(ulid.Timestamp(t), entropy)
		fmt.Println("Ulid:", id)
	}
}

// BenchmarkUlidFixedEntropyString "github.com/oklog/ulid" with fixed entropy
func BenchmarkUlidFixedEntropyString(b *testing.B) {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = ulid.MustNew(ulid.Timestamp(t), entropy).String()
	}
}

// BenchmarkUlidRandomEverytime "github.com/oklog/ulid"
func BenchmarkUlidRandomEverytime(b *testing.B) {
	t := time.Now().UTC()
	src := rand.NewSource(t.UnixNano())

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		entropy := rand.New(src)
		_ = ulid.MustNew(ulid.Timestamp(t), entropy).String()
	}
}

// BenchmarkSonyflake "github.com/sony/sonyflake"
func BenchmarkSonyflake(b *testing.B) {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = flake.NextID()
	}
}

// BenchmarkSid "github.com/chilts/sid"
func BenchmarkSid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = sid.Id()
	}
}

// BenchmarkUUIDv4RFC4122String "github.com/gofrs/uuid"
func BenchmarkUUIDv4RFC4122String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		id, _ := uuid.NewV4()
		_ = id.String()
	}
}

// BenchmarkUUIDv4Raw "github.com/gofrs/uuid"
func BenchmarkUUIDv4Raw(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		uuid.NewV4()
	}
}

// BenchmarkGONanoID "github.com/matoous/go-nanoid"
func BenchmarkGONanoID(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = gonanoid.Nanoid()
	}
}

// BenchmarkGONanoIDCustom "github.com/matoous/go-nanoid"
func BenchmarkGONanoIDCustom(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = gonanoid.Generate("abcde", 64)
	}
}

// BenchmarkSnowflake "github.com/bwmarrin/snowflake"
func BenchmarkSnowflake(b *testing.B) {
	b.ReportAllocs()
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < b.N; i++ {
		_ = node.Generate()
	}
}

func BenchmarkRand(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = rand.Int63()
	}
}

func BenchmarkCryptoRand(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		nBig, err := crypto_rand.Int(crypto_rand.Reader, big.NewInt(math.MaxInt64))
		if err != nil {
			b.FailNow()
		}
		_ = nBig.Int64()
	}
}

func BenchmarkTimestamp(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = time.Now().Nanosecond()
	}
}
