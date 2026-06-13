# internal

В этой директории размещается код внутренних модулей приложения. Код внутри этого пакета недоступен для импорта в других приложениях.

Директория `internal/` является специальной в Go и обеспечивает инкапсуляцию кода на уровне модуля. Компилятор Go запрещает импорт пакетов из `internal/` за пределами родительского модуля.

## Profile diff
File: main
Type: cpu
Time: 2026-06-13 12:21:33 MSK
Duration: 60.29s, Total samples = 39.55s (65.60%)
Showing nodes accounting for 0.08s, 0.2% of 39.55s total
Dropped 15 nodes (cum <= 0.20s)
      flat  flat%   sum%        cum   cum%
    -0.40s  1.01%  1.01%     -0.05s  0.13%  crypto/internal/fips140/pbkdf2.Key[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }]
     0.22s  0.56%  0.46%      0.22s  0.56%  runtime.kevent
    -0.18s  0.46%  0.91%     -0.19s  0.48%  syscall.syscall
     0.15s  0.38%  0.53%      0.19s  0.48%  crypto/internal/fips140/sha256.(*Digest).Write
     0.13s  0.33%   0.2%      0.13s  0.33%  runtime.pthread_cond_signal
    -0.12s   0.3%  0.51%     -0.12s   0.3%  runtime.usleep
     0.11s  0.28%  0.23%      0.14s  0.35%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
    -0.10s  0.25%  0.48%      0.28s  0.71%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.10s  0.25%  0.23%      0.10s  0.25%  runtime.memmove
    -0.10s  0.25%  0.48%     -0.10s  0.25%  syscall.rawSyscall
     0.06s  0.15%  0.33%      0.06s  0.15%  runtime.memclrNoHeapPointers
     0.05s  0.13%   0.2%      0.05s  0.13%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.05s  0.13% 0.076%      0.05s  0.13%  runtime.pthread_cond_wait
    -0.04s   0.1%  0.18%     -0.04s   0.1%  compress/flate.(*huffmanEncoder).bitLength (inline)
    -0.04s   0.1%  0.28%     -0.06s  0.15%  runtime.mallocgcSmallScanNoHeader
    -0.03s 0.076%  0.35%      0.05s  0.13%  compress/flate.(*compressor).deflate
     0.03s 0.076%  0.28%      0.04s   0.1%  compress/flate.(*huffmanEncoder).bitCounts
     0.03s 0.076%   0.2%      0.18s  0.46%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.03s 0.076%  0.13%      0.03s 0.076%  internal/byteorder.BEUint32 (inline)
     0.03s 0.076% 0.051%      0.04s   0.1%  internal/runtime/maps.(*Map).getWithoutKeySmallFastStr
    -0.03s 0.076%  0.13%     -0.03s 0.076%  runtime.pthread_mutex_unlock
     0.02s 0.051% 0.076%      0.02s 0.051%  bytes.Join
    -0.02s 0.051%  0.13%     -0.04s   0.1%  crypto/internal/fips140/sha256.block
    -0.02s 0.051%  0.18%     -0.02s 0.051%  crypto/internal/fips140/sha256.blockSHA2
     0.02s 0.051%  0.13%      0.02s 0.051%  encoding/base64.(*Encoding).Decode
     0.02s 0.051% 0.076%      0.02s 0.051%  encoding/json.stateBeginValue
     0.02s 0.051% 0.025%      0.02s 0.051%  github.com/jackc/pgx/v5/pgtype.(*Map).planEncode
     0.02s 0.051% 0.025%     -0.08s   0.2%  net.(*sysDialer).dialSerial
     0.02s 0.051% 0.076%      0.02s 0.051%  net/http.validCookieValueByte
    -0.02s 0.051% 0.025%     -0.02s 0.051%  net/textproto.CanonicalMIMEHeaderKey
     0.02s 0.051% 0.076%      0.02s 0.051%  net/url.parseQuery
     0.02s 0.051%  0.13%      0.02s 0.051%  runtime.duffcopy
     0.02s 0.051%  0.18%      0.02s 0.051%  runtime.madvise
    -0.02s 0.051%  0.13%     -0.02s 0.051%  runtime.nanotime1
    -0.02s 0.051% 0.076%     -0.02s 0.051%  runtime.spanOf (inline)
     0.02s 0.051%  0.13%      0.01s 0.025%  runtime.step
     0.02s 0.051%  0.18%      0.03s 0.076%  sort.choosePivot
    -0.01s 0.025%  0.15%     -0.01s 0.025%  bufio.(*Reader).ReadSlice
     0.01s 0.025%  0.18%      0.01s 0.025%  bufio.(*Writer).WriteString
     0.01s 0.025%   0.2%      0.01s 0.025%  cmpbody
     0.01s 0.025%  0.23%      0.02s 0.051%  compress/flate.(*compressor).findMatch
     0.01s 0.025%  0.25%      0.01s 0.025%  compress/flate.(*compressor).reset
     0.01s 0.025%  0.28%      0.01s 0.025%  compress/flate.(*huffmanBitWriter).generateCodegen
     0.01s 0.025%   0.3%      0.08s   0.2%  compress/flate.(*huffmanEncoder).generate
     0.01s 0.025%  0.33%      0.01s 0.025%  compress/flate.hash4 (inline)
     0.01s 0.025%  0.35%      0.01s 0.025%  compress/flate.matchLen (inline)
    -0.01s 0.025%  0.33%     -0.01s 0.025%  compress/gzip.(*Writer).Write
     0.01s 0.025%  0.35%      0.01s 0.025%  context.(*cancelCtx).Done
     0.01s 0.025%  0.38%      0.01s 0.025%  context.(*cancelCtx).cancel
    -0.01s 0.025%  0.35%     -0.01s 0.025%  context.(*cancelCtx).propagateCancel
    -0.01s 0.025%  0.33%     -0.02s 0.051%  context.WithDeadlineCause
    -0.01s 0.025%   0.3%     -0.01s 0.025%  context.value
     0.01s 0.025%  0.33%      0.01s 0.025%  crypto/hmac.New.UnwrapNew[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }].func1
    -0.01s 0.025%   0.3%     -0.02s 0.051%  crypto/internal/fips140.RecordApproved
    -0.01s 0.025%  0.28%     -0.01s 0.025%  crypto/internal/fips140.getIndicator
     0.01s 0.025%   0.3%      0.19s  0.48%  crypto/internal/fips140/sha256.(*Digest).Sum
    -0.01s 0.025%  0.28%     -0.01s 0.025%  encoding/hex.Encode (inline)
    -0.01s 0.025%  0.25%     -0.01s 0.025%  encoding/json.(*decodeState).literalStore
     0.01s 0.025%  0.28%     -0.03s 0.076%  encoding/json.(*decodeState).object
    -0.01s 0.025%  0.25%     -0.01s 0.025%  encoding/json.(*decodeState).scanWhile
     0.01s 0.025%  0.28%     -0.01s 0.025%  encoding/json.(*decodeState).unmarshal
     0.01s 0.025%   0.3%      0.03s 0.076%  encoding/json.(*encodeState).reflectValue
     0.01s 0.025%  0.33%      0.01s 0.025%  encoding/json.appendString[go.shape.[]uint8]
     0.01s 0.025%  0.35%      0.04s   0.1%  encoding/json.checkValid
     0.01s 0.025%  0.38%      0.03s 0.076%  encoding/json.marshalerEncoder
    -0.01s 0.025%  0.35%     -0.01s 0.025%  encoding/json.state1
     0.01s 0.025%  0.38%      0.01s 0.025%  encoding/json.stateBeginString
     0.01s 0.025%   0.4%      0.01s 0.025%  encoding/json.unquoteBytes
    -0.01s 0.025%  0.38%     -0.01s 0.025%  fmt.(*pp).handleMethods
    -0.01s 0.025%  0.35%     -0.03s 0.076%  fmt.(*pp).printArg
    -0.01s 0.025%  0.33%     -1.83s  4.63%  fmt.Fprintf
     0.01s 0.025%  0.35%      0.05s  0.13%  github.com/golang-jwt/jwt/v4.(*Token).SignedString
     0.01s 0.025%  0.38%      0.18s  0.46%  github.com/google/uuid.NewRandom
    -0.01s 0.025%  0.35%     -0.01s 0.025%  github.com/google/uuid.ParseBytes
     0.01s 0.025%  0.38%     -0.03s 0.076%  github.com/jackc/pgx/v5/pgconn.(*PgConn).execExtendedSuffix
    -0.01s 0.025%  0.35%     -0.05s  0.13%  github.com/jackc/pgx/v5/pgconn.(*scramClient).clientFinalMessage
    -0.01s 0.025%  0.33%     -0.01s 0.025%  github.com/jackc/pgx/v5/pgtype.TryWrapBuiltinTypeScanPlan
     0.01s 0.025%  0.35%      0.01s 0.025%  github.com/jackc/pgx/v5/pgtype.TryWrapDerefPointerEncodePlan
    -0.01s 0.025%  0.33%     -0.01s 0.025%  github.com/jackc/pgx/v5/pgtype.TryWrapPtrArrayScanPlan
     0.01s 0.025%  0.35%      0.01s 0.025%  github.com/labstack/echo/v4.(*Router).Find
    -0.01s 0.025%  0.33%     -0.01s 0.025%  github.com/labstack/echo/v4/middleware.GzipWithConfig.bufferPool.func3
     0.01s 0.025%  0.35%      0.01s 0.025%  github.com/labstack/echo/v4/middleware.GzipWithConfig.func1
    -0.01s 0.025%  0.33%     -0.01s 0.025%  go.uber.org/zap.Duration
    -0.01s 0.025%   0.3%     -0.01s 0.025%  go.uber.org/zap.Int
     0.01s 0.025%  0.33%      0.01s 0.025%  golang.org/x/text/secure/precis.(*buffers).apply
     0.01s 0.025%  0.35%      0.01s 0.025%  internal/abi.Name.ReadVarint (inline)
     0.01s 0.025%  0.38%      0.01s 0.025%  internal/byteorder.BEPutUint32 (inline)
     0.01s 0.025%   0.4%      0.01s 0.025%  internal/byteorder.BEPutUint64 (inline)
    -0.01s 0.025%  0.38%     -0.09s  0.23%  internal/poll.(*FD).Read
     0.01s 0.025%   0.4%      0.01s 0.025%  internal/runtime/atomic.(*Int32).Load (inline)
     0.01s 0.025%  0.43%      0.01s 0.025%  internal/runtime/atomic.(*Int32).Swap (inline)
    -0.01s 0.025%   0.4%     -0.01s 0.025%  internal/runtime/atomic.(*Uint32).CompareAndSwap (inline)
    -0.01s 0.025%  0.38%     -0.01s 0.025%  internal/runtime/atomic.(*Uintptr).Swap (inline)
    -0.01s 0.025%  0.35%     -0.01s 0.025%  internal/runtime/maps.(*Iter).Init
     0.01s 0.025%  0.38%      0.02s 0.051%  internal/runtime/maps.(*Map).Delete
    -0.01s 0.025%  0.35%      0.01s 0.025%  internal/runtime/maps.(*Map).growToSmall
     0.01s 0.025%  0.38%     -0.01s 0.025%  internal/runtime/maps.(*Map).putSlotSmall
    -0.01s 0.025%  0.35%     -0.01s 0.025%  internal/runtime/maps.(*groupReference).elem (inline)
    -0.01s 0.025%  0.33%     -0.01s 0.025%  internal/runtime/maps.(*groupsReference).group (inline)
     0.01s 0.025%  0.35%      0.01s 0.025%  internal/runtime/maps.typedmemclr
     0.01s 0.025%  0.38%      0.01s 0.025%  internal/sync.(*HashTrieMap[go.shape.interface {},go.shape.interface {}]).Load
     0.01s 0.025%   0.4%      0.01s 0.025%  internal/sync.(*Mutex).Lock (inline)
    -0.01s 0.025%  0.38%     -0.01s 0.025%  net.favoriteAddrFamily
    -0.01s 0.025%  0.35%     -0.01s 0.025%  net/http.(*body).Read
     0.01s 0.025%  0.38%      0.01s 0.025%  net/http.(*chunkWriter).close
     0.01s 0.025%   0.4%      0.01s 0.025%  net/http.(*transferReader).parseTransferEncoding
    -0.01s 0.025%  0.38%      0.01s 0.025%  net/http.sanitizeOrWarn
     0.01s 0.025%   0.4%     -0.15s  0.38%  net/http.serverHandler.ServeHTTP
     0.01s 0.025%  0.43%      0.01s 0.025%  net/textproto.TrimString (inline)
    -0.01s 0.025%   0.4%     -0.01s 0.025%  net/url.shouldEscape
    -0.01s 0.025%  0.38%     -0.01s 0.025%  reflect.(*rtype).typeOff (inline)
    -0.01s 0.025%  0.35%     -0.01s 0.025%  reflect.TypeOf (inline)
    -0.01s 0.025%  0.33%     -0.01s 0.025%  reflect.Value.IsNil (inline)
    -0.01s 0.025%   0.3%     -0.01s 0.025%  reflect.maplen
     0.01s 0.025%  0.33%      0.01s 0.025%  reflect.packEface
     0.01s 0.025%  0.35%      0.01s 0.025%  reflect.unpackEface (inline)
     0.01s 0.025%  0.38%      0.01s 0.025%  runtime.(*gcBitsArena).tryAlloc (inline)
    -0.01s 0.025%  0.35%     -0.01s 0.025%  runtime.(*gcControllerState).update
    -0.01s 0.025%  0.33%      0.03s 0.076%  runtime.(*mcentral).cacheSpan
    -0.01s 0.025%   0.3%      0.05s  0.13%  runtime.(*mheap).alloc.func1
     0.01s 0.025%  0.33%      0.01s 0.025%  runtime.(*mspan).heapBitsSmallForAddr
     0.01s 0.025%  0.35%      0.01s 0.025%  runtime.(*mspan).init
     0.01s 0.025%  0.38%      0.01s 0.025%  runtime.(*pallocBits).summarize
    -0.01s 0.025%  0.35%     -0.01s 0.025%  runtime.(*randomOrder).start (inline)
    -0.01s 0.025%  0.33%     -0.01s 0.025%  runtime.(*semaRoot).dequeue
    -0.01s 0.025%   0.3%     -0.02s 0.051%  runtime.(*timers).check
    -0.01s 0.025%  0.28%     -0.01s 0.025%  runtime.(*timers).wakeTime (inline)
     0.01s 0.025%   0.3%      0.01s 0.025%  runtime.(*unwinder).resolveInternal
     0.01s 0.025%  0.33%      0.01s 0.025%  runtime.SetFinalizer
    -0.01s 0.025%   0.3%     -0.01s 0.025%  runtime.acquirem (inline)
     0.01s 0.025%  0.33%      0.02s 0.051%  runtime.concatstring5
     0.01s 0.025%  0.35%      0.01s 0.025%  runtime.concatstrings
     0.01s 0.025%  0.38%      0.01s 0.025%  runtime.convTnoptr
     0.01s 0.025%   0.4%      0.01s 0.025%  runtime.divRoundUp (inline)
    -0.01s 0.025%  0.38%     -0.02s 0.051%  runtime.execute
     0.01s 0.025%   0.4%      0.01s 0.025%  runtime.exitsyscallfast
    -0.01s 0.025%  0.38%     -0.01s 0.025%  runtime.getMCache (inline)
     0.01s 0.025%   0.4%      0.01s 0.025%  runtime.gfget
    -0.01s 0.025%  0.38%     -0.01s 0.025%  runtime.heapSetTypeNoHeader (inline)
     0.01s 0.025%   0.4%     -0.06s  0.15%  runtime.injectglist.func1
     0.01s 0.025%  0.43%     -0.12s   0.3%  runtime.lock2
     0.01s 0.025%  0.46%      0.01s 0.025%  runtime.makeSpanClass (inline)
     0.01s 0.025%  0.48%      0.04s   0.1%  runtime.mallocgcSmallNoscan
     0.01s 0.025%  0.51%      0.01s 0.025%  runtime.mallocgcTiny
     0.01s 0.025%  0.53%      0.01s 0.025%  runtime.mapaccess1
    -0.01s 0.025%  0.51%     -0.01s 0.025%  runtime.mapaccess1_fast32
    -0.01s 0.025%  0.48%     -0.01s 0.025%  runtime.memequal
    -0.01s 0.025%  0.46%      0.21s  0.53%  runtime.netpoll
    -0.01s 0.025%  0.43%     -0.01s 0.025%  runtime.netpollunblock (inline)
     0.01s 0.025%  0.46%      0.02s 0.051%  runtime.newMarkBits
    -0.01s 0.025%  0.43%     -0.01s 0.025%  runtime.nextFreeFast (inline)
     0.01s 0.025%  0.46%      0.01s 0.025%  runtime.pMask.clear (inline)
    -0.01s 0.025%  0.43%     -0.01s 0.025%  runtime.pthread_kill
    -0.01s 0.025%   0.4%     -0.01s 0.025%  runtime.readvarint (inline)
     0.01s 0.025%  0.43%      0.01s 0.025%  runtime.releasem (inline)
    -0.01s 0.025%   0.4%     -0.01s 0.025%  runtime.releasep
    -0.01s 0.025%  0.38%     -0.01s 0.025%  runtime.runqempty (inline)
    -0.01s 0.025%  0.35%     -0.01s 0.025%  runtime.save
     0.01s 0.025%  0.38%      0.01s 0.025%  runtime.save_g
    -0.01s 0.025%  0.35%     -0.01s 0.025%  runtime.scanblock
    -0.01s 0.025%  0.33%     -0.23s  0.58%  runtime.schedule
    -0.01s 0.025%   0.3%     -0.01s 0.025%  runtime.setprofilebucket
     0.01s 0.025%  0.33%      0.01s 0.025%  runtime.strhash
     0.01s 0.025%  0.35%      0.15s  0.38%  runtime.systemstack
    -0.01s 0.025%  0.33%     -0.01s 0.025%  runtime.taggedPointerPack
    -0.01s 0.025%   0.3%     -0.01s 0.025%  runtime.typedmemmove
    -0.01s 0.025%  0.28%     -0.01s 0.025%  runtime.wirep
    -0.01s 0.025%  0.25%     -0.01s 0.025%  slices.pdqsortCmpFunc[go.shape.struct { net/http.key string; net/http.values []string }]
     0.01s 0.025%  0.28%      0.01s 0.025%  sort.insertionSort
    -0.01s 0.025%  0.25%      0.03s 0.076%  sort.pdqsort
    -0.01s 0.025%  0.23%     -0.01s 0.025%  strconv.readFloat
     0.01s 0.025%  0.25%      0.01s 0.025%  strings.(*Builder).copyCheck (inline)
    -0.01s 0.025%  0.23%     -0.01s 0.025%  strings.(*byteReplacer).Replace
    -0.01s 0.025%   0.2%     -0.01s 0.025%  sync.(*Once).doSlow
     0.01s 0.025%  0.23%     -0.01s 0.025%  sync.(*Pool).Put
    -0.01s 0.025%   0.2%     -0.01s 0.025%  sync.(*poolDequeue).pushHead
    -0.01s 0.025%  0.18%     -0.01s 0.025%  time.absDays.date
     0.01s 0.025%   0.2%      0.01s 0.025%  time.appendInt
    -0.01s 0.025%  0.18%     -0.01s 0.025%  time.nextStdChunk
     0.01s 0.025%   0.2%      0.01s 0.025%  time.now
    -0.01s 0.025%  0.18%     -0.01s 0.025%  unicode/utf8.AppendRune (inline)
     0.01s 0.025%   0.2%      0.01s 0.025%  vendor/golang.org/x/net/http/httpguts.ValidHeaderFieldName (inline)
         0     0%   0.2%      0.10s  0.25%  bufio.(*Reader).Peek
         0     0%   0.2%     -0.01s 0.025%  bufio.(*Reader).ReadLine
         0     0%   0.2%      0.10s  0.25%  bufio.(*Reader).fill
         0     0%   0.2%      0.21s  0.53%  bufio.(*Writer).Flush
         0     0%   0.2%     -0.01s 0.025%  bufio.(*Writer).Write
         0     0%   0.2%      0.05s  0.13%  compress/flate.(*Writer).Close (inline)
         0     0%   0.2%      0.01s 0.025%  compress/flate.(*Writer).Reset
         0     0%   0.2%      0.03s 0.076%  compress/flate.(*byFreq).sort (inline)
         0     0%   0.2%      0.05s  0.13%  compress/flate.(*compressor).close
         0     0%   0.2%      0.05s  0.13%  compress/flate.(*compressor).writeBlock
         0     0%   0.2%     -0.01s 0.025%  compress/flate.(*huffmanBitWriter).dynamicSize
         0     0%   0.2%     -0.03s 0.076%  compress/flate.(*huffmanBitWriter).fixedSize (inline)
         0     0%   0.2%      0.07s  0.18%  compress/flate.(*huffmanBitWriter).indexTokens
         0     0%   0.2%      0.05s  0.13%  compress/flate.(*huffmanBitWriter).writeBlock
         0     0%   0.2%      0.05s  0.13%  compress/gzip.(*Writer).Close
         0     0%   0.2%      0.01s 0.025%  compress/gzip.(*Writer).Reset (inline)
         0     0%   0.2%      0.01s 0.025%  compress/gzip.(*Writer).init
         0     0%   0.2%     -0.01s 0.025%  context.(*afterFuncCtx).cancel
         0     0%   0.2%     -0.01s 0.025%  context.(*cancelCtx).Value
         0     0%   0.2%     -0.02s 0.051%  context.AfterFunc
         0     0%   0.2%     -0.02s 0.051%  context.AfterFunc.func1
         0     0%   0.2%      0.02s 0.051%  context.WithCancel.func1
         0     0%   0.2%     -0.02s 0.051%  context.WithDeadline (inline)
         0     0%   0.2%     -0.01s 0.025%  context.WithTimeout
         0     0%   0.2%      0.01s 0.025%  context.removeChild
         0     0%   0.2%     -0.01s 0.025%  crypto.Hash.New
         0     0%   0.2%      0.02s 0.051%  crypto/hmac.New
         0     0%   0.2%      0.02s 0.051%  crypto/internal/fips140/hmac.New[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }]
         0     0%   0.2%     -0.01s 0.025%  crypto/internal/fips140/hmac.New[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }].func1
         0     0%   0.2%      0.03s 0.076%  crypto/internal/fips140/sha256.consumeUint32 (inline)
         0     0%   0.2%      0.01s 0.025%  crypto/internal/fips140deps/byteorder.BEPutUint32 (inline)
         0     0%   0.2%      0.01s 0.025%  crypto/internal/fips140deps/byteorder.BEPutUint64 (inline)
         0     0%   0.2%      0.03s 0.076%  crypto/internal/fips140deps/byteorder.BEUint32 (inline)
         0     0%   0.2%      0.18s  0.46%  crypto/internal/sysrand.Read
         0     0%   0.2%      0.18s  0.46%  crypto/internal/sysrand.read (inline)
         0     0%   0.2%     -0.05s  0.13%  crypto/pbkdf2.Key[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }]
         0     0%   0.2%      0.18s  0.46%  crypto/rand.(*reader).Read
         0     0%   0.2%      0.01s 0.025%  crypto/rand.Read
         0     0%   0.2%      1.33s  3.36%  database/sql.(*DB).BeginTx
         0     0%   0.2%      1.33s  3.36%  database/sql.(*DB).BeginTx.func1
         0     0%   0.2%      0.01s 0.025%  database/sql.(*DB).addDepLocked (inline)
         0     0%   0.2%      1.33s  3.36%  database/sql.(*DB).begin
         0     0%   0.2%      0.69s  1.74%  database/sql.(*DB).beginDC
         0     0%   0.2%      0.68s  1.72%  database/sql.(*DB).beginDC.func1
         0     0%   0.2%      0.64s  1.62%  database/sql.(*DB).conn
         0     0%   0.2%     -0.03s 0.076%  database/sql.(*DB).prepareDC
         0     0%   0.2%     -0.03s 0.076%  database/sql.(*DB).prepareDC.func2
         0     0%   0.2%      0.01s 0.025%  database/sql.(*DB).putConn
         0     0%   0.2%      1.26s  3.19%  database/sql.(*DB).retry
         0     0%   0.2%      0.09s  0.23%  database/sql.(*Stmt).Close
         0     0%   0.2%     -0.07s  0.18%  database/sql.(*Stmt).ExecContext
         0     0%   0.2%     -0.07s  0.18%  database/sql.(*Stmt).ExecContext.func1
         0     0%   0.2%      0.14s  0.35%  database/sql.(*Tx).Commit
         0     0%   0.2%      0.04s   0.1%  database/sql.(*Tx).Commit.func1
         0     0%   0.2%     -0.03s 0.076%  database/sql.(*Tx).PrepareContext
         0     0%   0.2%      0.01s 0.025%  database/sql.(*Tx).close (inline)
         0     0%   0.2%      0.09s  0.23%  database/sql.(*Tx).closePrepared
         0     0%   0.2%      0.01s 0.025%  database/sql.(*driverConn).Close
         0     0%   0.2%      0.01s 0.025%  database/sql.(*driverConn).finalClose
         0     0%   0.2%      0.01s 0.025%  database/sql.(*driverConn).finalClose.func2
         0     0%   0.2%     -0.03s 0.076%  database/sql.(*driverConn).prepareLocked
         0     0%   0.2%      0.01s 0.025%  database/sql.(*driverConn).releaseConn
         0     0%   0.2%      0.24s  0.61%  database/sql.(*driverConn).resetSession
         0     0%   0.2%      0.09s  0.23%  database/sql.(*driverStmt).Close
         0     0%   0.2%      0.68s  1.72%  database/sql.ctxDriverBegin
         0     0%   0.2%     -0.03s 0.076%  database/sql.ctxDriverPrepare
         0     0%   0.2%     -0.06s  0.15%  database/sql.ctxDriverStmtExec
         0     0%   0.2%     -0.01s 0.025%  database/sql.driverArgsConnLocked
         0     0%   0.2%     -0.07s  0.18%  database/sql.resultFromStatement
         0     0%   0.2%      0.70s  1.77%  database/sql.withLock
         0     0%   0.2%      0.01s 0.025%  encoding/base64.(*Encoding).DecodeString
         0     0%   0.2%     -0.05s  0.13%  encoding/json.(*Decoder).Decode
         0     0%   0.2%     -0.01s 0.025%  encoding/json.(*Decoder).readValue
         0     0%   0.2%     -0.03s 0.076%  encoding/json.(*Encoder).Encode
         0     0%   0.2%      0.04s   0.1%  encoding/json.(*decodeState).array
         0     0%   0.2%     -0.02s 0.051%  encoding/json.(*decodeState).value
         0     0%   0.2%      0.03s 0.076%  encoding/json.(*encodeState).marshal
         0     0%   0.2%      0.04s   0.1%  encoding/json.Marshal
         0     0%   0.2%     -0.01s 0.025%  encoding/json.Number.Float64 (inline)
         0     0%   0.2%      0.06s  0.15%  encoding/json.Unmarshal
         0     0%   0.2%     -0.01s 0.025%  encoding/json.arrayEncoder.encode
         0     0%   0.2%      0.01s 0.025%  encoding/json.condAddrEncoder.encode
         0     0%   0.2%     -0.01s 0.025%  encoding/json.interfaceEncoder
         0     0%   0.2%     -0.01s 0.025%  encoding/json.isEmptyValue
         0     0%   0.2%     -0.02s 0.051%  encoding/json.mapEncoder.encode
         0     0%   0.2%      0.01s 0.025%  encoding/json.newEncodeState
         0     0%   0.2%     -0.01s 0.025%  encoding/json.sliceEncoder.encode
         0     0%   0.2%      0.03s 0.076%  encoding/json.structEncoder.encode
         0     0%   0.2%      0.01s 0.025%  encoding/json.textMarshalerEncoder
         0     0%   0.2%      0.01s 0.025%  encoding/json.typeEncoder
         0     0%   0.2%      0.01s 0.025%  encoding/json.valueEncoder
         0     0%   0.2%     -0.01s 0.025%  fmt.(*buffer).writeString (inline)
         0     0%   0.2%     -0.01s 0.025%  fmt.(*fmt).fmtQ
         0     0%   0.2%     -0.04s   0.1%  fmt.(*pp).doPrintf
         0     0%   0.2%     -0.01s 0.025%  fmt.(*pp).fmtString
         0     0%   0.2%     -0.04s   0.1%  fmt.Errorf
         0     0%   0.2%     -1.83s  4.63%  fmt.Printf (inline)
         0     0%   0.2%     -0.02s 0.051%  github.com/golang-jwt/jwt/v4.(*NumericDate).UnmarshalJSON
         0     0%   0.2%     -0.04s   0.1%  github.com/golang-jwt/jwt/v4.(*Parser).ParseUnverified
         0     0%   0.2%     -0.04s   0.1%  github.com/golang-jwt/jwt/v4.(*Parser).ParseWithClaims
         0     0%   0.2%     -0.01s 0.025%  github.com/golang-jwt/jwt/v4.(*SigningMethodHMAC).Sign
         0     0%   0.2%      0.05s  0.13%  github.com/golang-jwt/jwt/v4.(*Token).SigningString
         0     0%   0.2%      0.01s 0.025%  github.com/golang-jwt/jwt/v4.DecodeSegment
         0     0%   0.2%      0.01s 0.025%  github.com/golang-jwt/jwt/v4.NewWithClaims (inline)
         0     0%   0.2%      0.01s 0.025%  github.com/golang-jwt/jwt/v4.NumericDate.MarshalJSON
         0     0%   0.2%     -0.04s   0.1%  github.com/golang-jwt/jwt/v4.ParseWithClaims
         0     0%   0.2%     -0.01s 0.025%  github.com/google/uuid.(*UUID).UnmarshalText
         0     0%   0.2%      0.18s  0.46%  github.com/google/uuid.New (inline)
         0     0%   0.2%      0.17s  0.43%  github.com/google/uuid.NewRandomFromReader
         0     0%   0.2%     -0.01s 0.025%  github.com/google/uuid.UUID.Value
         0     0%   0.2%     -0.01s 0.025%  github.com/google/uuid.encodeHex
         0     0%   0.2%     -0.16s   0.4%  github.com/grizlaz/ya-shortener/internal/handler.(*Server).ServeHTTP
         0     0%   0.2%     -0.20s  0.51%  github.com/grizlaz/ya-shortener/internal/handler.(*Server).setupRouter.DecompressWithConfig.func7.1
         0     0%   0.2%     -0.29s  0.73%  github.com/grizlaz/ya-shortener/internal/handler.(*Server).setupRouter.HandleAPIShortenBatch.func12
         0     0%   0.2%     -0.15s  0.38%  github.com/grizlaz/ya-shortener/internal/handler.(*Server).setupRouter.RemoveTrailingSlash.RemoveTrailingSlashWithConfig.func17.1
         0     0%   0.2%     -0.19s  0.48%  github.com/grizlaz/ya-shortener/internal/handler.(*Server).setupRouter.WithJWT.func9.1
         0     0%   0.2%     -0.20s  0.51%  github.com/grizlaz/ya-shortener/internal/handler.(*Server).setupRouter.WithLogging.func8.1
         0     0%   0.2%      0.01s 0.025%  github.com/grizlaz/ya-shortener/internal/handler.getUserID
         0     0%   0.2%      0.06s  0.15%  github.com/grizlaz/ya-shortener/internal/handler.makeJWT
         0     0%   0.2%      1.37s  3.46%  github.com/grizlaz/ya-shortener/internal/repository.(*postgres).PutBatch
         0     0%   0.2%     -0.33s  0.83%  github.com/grizlaz/ya-shortener/internal/service.(*Service).ShortenBatch
         0     0%   0.2%      0.01s 0.025%  github.com/grizlaz/ya-shortener/internal/service.PrependBaseURL
         0     0%   0.2%      0.09s  0.23%  github.com/jackc/pgpassfile.ReadPassfile
         0     0%   0.2%      0.68s  1.72%  github.com/jackc/pgx/v5.(*Conn).BeginTx
         0     0%   0.2%      0.01s 0.025%  github.com/jackc/pgx/v5.(*Conn).Close
         0     0%   0.2%      0.10s  0.25%  github.com/jackc/pgx/v5.(*Conn).Deallocate
         0     0%   0.2%      0.66s  1.67%  github.com/jackc/pgx/v5.(*Conn).Exec
         0     0%   0.2%     -0.03s 0.076%  github.com/jackc/pgx/v5.(*Conn).Prepare
         0     0%   0.2%      0.66s  1.67%  github.com/jackc/pgx/v5.(*Conn).exec
         0     0%   0.2%     -0.06s  0.15%  github.com/jackc/pgx/v5.(*Conn).execPrepared
         0     0%   0.2%      0.72s  1.82%  github.com/jackc/pgx/v5.(*Conn).execSimpleProtocol
         0     0%   0.2%     -0.01s 0.025%  github.com/jackc/pgx/v5.(*ConnConfig).Copy (inline)
         0     0%   0.2%     -0.02s 0.051%  github.com/jackc/pgx/v5.(*ExtendedQueryBuilder).Build
         0     0%   0.2%     -0.02s 0.051%  github.com/jackc/pgx/v5.(*ExtendedQueryBuilder).appendParam
         0     0%   0.2%     -0.02s 0.051%  github.com/jackc/pgx/v5.(*ExtendedQueryBuilder).encodeExtendedParamValue
         0     0%   0.2%      0.04s   0.1%  github.com/jackc/pgx/v5.(*dbTx).Commit
         0     0%   0.2%      0.04s   0.1%  github.com/jackc/pgx/v5.ConnectConfig
         0     0%   0.2%      0.36s  0.91%  github.com/jackc/pgx/v5.ParseConfig (inline)
         0     0%   0.2%      0.36s  0.91%  github.com/jackc/pgx/v5.ParseConfigWithOptions
         0     0%   0.2%      0.05s  0.13%  github.com/jackc/pgx/v5.connect
         0     0%   0.2%      0.01s 0.025%  github.com/jackc/pgx/v5/internal/iobufpool.Get
         0     0%   0.2%      0.01s 0.025%  github.com/jackc/pgx/v5/internal/iobufpool.init.0.func1
         0     0%   0.2%      0.04s   0.1%  github.com/jackc/pgx/v5/internal/stmtcache.NewLRUCache (inline)
         0     0%   0.2%     -0.01s 0.025%  github.com/jackc/pgx/v5/pgconn.(*Config).Copy
         0     0%   0.2%     -0.02s 0.051%  github.com/jackc/pgx/v5/pgconn.(*MultiResultReader).Close (inline)
         0     0%   0.2%     -0.04s   0.1%  github.com/jackc/pgx/v5/pgconn.(*MultiResultReader).NextResult
         0     0%   0.2%     -0.02s 0.051%  github.com/jackc/pgx/v5/pgconn.(*MultiResultReader).receiveMessage
         0     0%   0.2%      0.01s 0.025%  github.com/jackc/pgx/v5/pgconn.(*PgConn).Close
         0     0%   0.2%      0.10s  0.25%  github.com/jackc/pgx/v5/pgconn.(*PgConn).Deallocate
         0     0%   0.2%      1.02s  2.58%  github.com/jackc/pgx/v5/pgconn.(*PgConn).Exec
         0     0%   0.2%     -0.04s   0.1%  github.com/jackc/pgx/v5/pgconn.(*PgConn).ExecStatement
         0     0%   0.2%      0.24s  0.61%  github.com/jackc/pgx/v5/pgconn.(*PgConn).Ping
         0     0%   0.2%     -0.03s 0.076%  github.com/jackc/pgx/v5/pgconn.(*PgConn).Prepare
         0     0%   0.2%      0.37s  0.94%  github.com/jackc/pgx/v5/pgconn.(*PgConn).enterPotentialWriteReadDeadlock (inline)
         0     0%   0.2%     -0.01s 0.025%  github.com/jackc/pgx/v5/pgconn.(*PgConn).execExtendedPrefix
         0     0%   0.2%      1.27s  3.21%  github.com/jackc/pgx/v5/pgconn.(*PgConn).flushWithPotentialWriteReadDeadlock
         0     0%   0.2%     -0.11s  0.28%  github.com/jackc/pgx/v5/pgconn.(*PgConn).peekMessage
         0     0%   0.2%     -0.09s  0.23%  github.com/jackc/pgx/v5/pgconn.(*PgConn).receiveMessage
         0     0%   0.2%     -0.01s 0.025%  github.com/jackc/pgx/v5/pgconn.(*PgConn).rxSASLContinue
         0     0%   0.2%     -0.02s 0.051%  github.com/jackc/pgx/v5/pgconn.(*PgConn).rxSASLFinal
         0     0%   0.2%      0.26s  0.66%  github.com/jackc/pgx/v5/pgconn.(*PgConn).scramAuth
         0     0%   0.2%     -0.01s 0.025%  github.com/jackc/pgx/v5/pgconn.(*ResultReader).Close
         0     0%   0.2%     -0.07s  0.18%  github.com/jackc/pgx/v5/pgconn.(*ResultReader).readUntilRowDescription
         0     0%   0.2%     -0.07s  0.18%  github.com/jackc/pgx/v5/pgconn.(*ResultReader).receiveMessage
         0     0%   0.2%      0.01s 0.025%  github.com/jackc/pgx/v5/pgconn.ConnectConfig
         0     0%   0.2%      0.36s  0.91%  github.com/jackc/pgx/v5/pgconn.ParseConfigWithOptions
         0     0%   0.2%      0.01s 0.025%  github.com/jackc/pgx/v5/pgconn.ParseConfigWithOptions.func1
         0     0%   0.2%     -0.02s 0.051%  github.com/jackc/pgx/v5/pgconn.buildConnectOneConfigs
         0     0%   0.2%     -0.01s 0.025%  github.com/jackc/pgx/v5/pgconn.computeClientProof
         0     0%   0.2%      0.03s 0.076%  github.com/jackc/pgx/v5/pgconn.connectOne
         0     0%   0.2%      0.03s 0.076%  github.com/jackc/pgx/v5/pgconn.connectPreferred
         0     0%   0.2%      0.28s  0.71%  github.com/jackc/pgx/v5/pgconn.defaultHost
         0     0%   0.2%      0.27s  0.68%  github.com/jackc/pgx/v5/pgconn.defaultSettings
         0     0%   0.2%      0.02s 0.051%  github.com/jackc/pgx/v5/pgconn.newScramClient
         0     0%   0.2%     -0.02s 0.051%  github.com/jackc/pgx/v5/pgconn/ctxwatch.(*ContextWatcher).Unwatch
         0     0%   0.2%     -0.11s  0.28%  github.com/jackc/pgx/v5/pgconn/internal/bgreader.(*BGReader).Read
         0     0%   0.2%      0.90s  2.28%  github.com/jackc/pgx/v5/pgproto3.(*Frontend).Flush
         0     0%   0.2%     -0.11s  0.28%  github.com/jackc/pgx/v5/pgproto3.(*Frontend).Receive
         0     0%   0.2%     -0.11s  0.28%  github.com/jackc/pgx/v5/pgproto3.(*chunkReader).Next
         0     0%   0.2%      0.01s 0.025%  github.com/jackc/pgx/v5/pgproto3.NewFrontend
         0     0%   0.2%      0.01s 0.025%  github.com/jackc/pgx/v5/pgproto3.newChunkReader (inline)
         0     0%   0.2%     -0.02s 0.051%  github.com/jackc/pgx/v5/pgtype.(*Map).Encode
         0     0%   0.2%      0.01s 0.025%  github.com/jackc/pgx/v5/pgtype.(*Map).PlanEncode (inline)
         0     0%   0.2%     -0.03s 0.076%  github.com/jackc/pgx/v5/pgtype.(*Map).PlanScan (inline)
         0     0%   0.2%      0.01s 0.025%  github.com/jackc/pgx/v5/pgtype.(*Map).planEncodeDepth
         0     0%   0.2%     -0.03s 0.076%  github.com/jackc/pgx/v5/pgtype.(*Map).planScan
         0     0%   0.2%     -0.03s 0.076%  github.com/jackc/pgx/v5/pgtype.(*encodePlanDriverValuer).Encode
         0     0%   0.2%      0.03s 0.076%  github.com/jackc/pgx/v5/pgtype.(*pointerEmptyInterfaceScanPlan).Scan
         0     0%   0.2%      0.01s 0.025%  github.com/jackc/pgx/v5/pgtype.TryPointerPointerScanPlan
         0     0%   0.2%      0.03s 0.076%  github.com/jackc/pgx/v5/pgtype.UUIDCodec.DecodeValue
         0     0%   0.2%      0.02s 0.051%  github.com/jackc/pgx/v5/pgtype.codecScan
         0     0%   0.2%      0.01s 0.025%  github.com/jackc/pgx/v5/pgtype.isNilDriverValuer
         0     0%   0.2%     -0.04s   0.1%  github.com/jackc/pgx/v5/pgtype.newEncodeError
         0     0%   0.2%      0.02s 0.051%  github.com/jackc/pgx/v5/pgtype.parseUUID
         0     0%   0.2%      0.02s 0.051%  github.com/jackc/pgx/v5/pgtype.scanPlanTextAnyToUUIDScanner.Scan
         0     0%   0.2%      0.68s  1.72%  github.com/jackc/pgx/v5/stdlib.(*Conn).BeginTx
         0     0%   0.2%      0.01s 0.025%  github.com/jackc/pgx/v5/stdlib.(*Conn).Close
         0     0%   0.2%     -0.06s  0.15%  github.com/jackc/pgx/v5/stdlib.(*Conn).ExecContext
         0     0%   0.2%     -0.03s 0.076%  github.com/jackc/pgx/v5/stdlib.(*Conn).PrepareContext
         0     0%   0.2%      0.24s  0.61%  github.com/jackc/pgx/v5/stdlib.(*Conn).ResetSession
         0     0%   0.2%      0.09s  0.23%  github.com/jackc/pgx/v5/stdlib.(*Stmt).Close
         0     0%   0.2%     -0.06s  0.15%  github.com/jackc/pgx/v5/stdlib.(*Stmt).ExecContext
         0     0%   0.2%      0.39s  0.99%  github.com/jackc/pgx/v5/stdlib.(*driverConnector).Connect
         0     0%   0.2%      0.04s   0.1%  github.com/jackc/pgx/v5/stdlib.wrapTx.Commit
         0     0%   0.2%     -0.16s   0.4%  github.com/labstack/echo/v4.(*Echo).ServeHTTP
         0     0%   0.2%     -0.15s  0.38%  github.com/labstack/echo/v4.(*Echo).ServeHTTP.func1
         0     0%   0.2%     -0.29s  0.73%  github.com/labstack/echo/v4.(*Echo).add.func1
         0     0%   0.2%     -0.03s 0.076%  github.com/labstack/echo/v4.(*Response).Write
         0     0%   0.2%      0.01s 0.025%  github.com/labstack/echo/v4.(*context).Cookie
         0     0%   0.2%     -0.02s 0.051%  github.com/labstack/echo/v4.(*context).JSON
         0     0%   0.2%      0.02s 0.051%  github.com/labstack/echo/v4.(*context).QueryParams (inline)
         0     0%   0.2%     -0.04s   0.1%  github.com/labstack/echo/v4.(*context).json
         0     0%   0.2%     -0.01s 0.025%  github.com/labstack/echo/v4.(*context).writeContentType
         0     0%   0.2%     -0.03s 0.076%  github.com/labstack/echo/v4.DefaultJSONSerializer.Serialize
         0     0%   0.2%     -0.03s 0.076%  github.com/labstack/echo/v4/middleware.(*gzipResponseWriter).Write
         0     0%   0.2%     -0.16s   0.4%  github.com/labstack/echo/v4/middleware.GzipWithConfig.func1.1
         0     0%   0.2%      0.04s   0.1%  github.com/labstack/echo/v4/middleware.GzipWithConfig.func1.1.1
         0     0%   0.2%      0.01s 0.025%  go.uber.org/zap.(*Logger).Info
         0     0%   0.2%      0.01s 0.025%  go.uber.org/zap.(*Logger).check
         0     0%   0.2%      0.01s 0.025%  go.uber.org/zap/buffer.(*Buffer).AppendInt (inline)
         0     0%   0.2%      0.01s 0.025%  go.uber.org/zap/zapcore.(*jsonEncoder).EncodeEntry
         0     0%   0.2%     -0.01s 0.025%  go.uber.org/zap/zapcore.(*lockedWriteSyncer).Write
         0     0%   0.2%      0.01s 0.025%  go.uber.org/zap/zapcore.EntryCaller.TrimmedPath
         0     0%   0.2%      0.01s 0.025%  go.uber.org/zap/zapcore.ShortCallerEncoder
         0     0%   0.2%      0.01s 0.025%  go.uber.org/zap/zapcore.systemClock.Now
         0     0%   0.2%      0.01s 0.025%  golang.org/x/text/secure/precis.(*Profile).String (inline)
         0     0%   0.2%      0.01s 0.025%  golang.org/x/text/secure/precis.(*buffers).enforce
         0     0%   0.2%      0.01s 0.025%  golang.org/x/text/secure/precis.processString
         0     0%   0.2%      0.01s 0.025%  internal/abi.Name.Name
         0     0%   0.2%      0.01s 0.025%  internal/bytealg.CompareString (inline)
         0     0%   0.2%      0.02s 0.051%  internal/poll.(*FD).Close
         0     0%   0.2%     -0.02s 0.051%  internal/poll.(*FD).Init
         0     0%   0.2%     -0.01s 0.025%  internal/poll.(*FD).SetReadDeadline (inline)
         0     0%   0.2%     -0.03s 0.076%  internal/poll.(*FD).SetsockoptInt
         0     0%   0.2%     -0.72s  1.82%  internal/poll.(*FD).Write
         0     0%   0.2%      0.02s 0.051%  internal/poll.(*FD).decref
         0     0%   0.2%      0.02s 0.051%  internal/poll.(*FD).destroy
         0     0%   0.2%     -0.01s 0.025%  internal/poll.(*FD).writeLock (inline)
         0     0%   0.2%     -0.01s 0.025%  internal/poll.(*FD).writeUnlock
         0     0%   0.2%      0.02s 0.051%  internal/poll.(*SysFile).destroy (inline)
         0     0%   0.2%     -0.01s 0.025%  internal/poll.(*fdMutex).rwlock
         0     0%   0.2%     -0.01s 0.025%  internal/poll.(*fdMutex).rwunlock
         0     0%   0.2%     -0.02s 0.051%  internal/poll.(*pollDesc).init
         0     0%   0.2%     -0.01s 0.025%  internal/poll.(*pollDesc).wait
         0     0%   0.2%     -0.01s 0.025%  internal/poll.(*pollDesc).waitRead (inline)
         0     0%   0.2%     -0.77s  1.95%  internal/poll.ignoringEINTRIO (inline)
         0     0%   0.2%     -0.01s 0.025%  internal/poll.runtime_Semacquire
         0     0%   0.2%     -0.01s 0.025%  internal/poll.runtime_Semrelease
         0     0%   0.2%     -0.02s 0.051%  internal/poll.runtime_pollOpen
         0     0%   0.2%     -0.01s 0.025%  internal/poll.runtime_pollWait
         0     0%   0.2%     -0.01s 0.025%  internal/poll.setDeadlineImpl
         0     0%   0.2%      0.01s 0.025%  internal/runtime/maps.(*Map).deleteSmall
         0     0%   0.2%     -0.01s 0.025%  internal/runtime/maps.(*table).grow
         0     0%   0.2%     -0.01s 0.025%  internal/runtime/maps.(*table).rehash
         0     0%   0.2%      0.04s   0.1%  internal/runtime/maps.(*table).reset
         0     0%   0.2%      0.01s 0.025%  internal/runtime/maps.NewEmptyMap (inline)
         0     0%   0.2%      0.03s 0.076%  internal/runtime/maps.NewMap
         0     0%   0.2%      0.07s  0.18%  internal/runtime/maps.newGroups (inline)
         0     0%   0.2%      0.04s   0.1%  internal/runtime/maps.newTable
         0     0%   0.2%      0.07s  0.18%  internal/runtime/maps.newarray
         0     0%   0.2%     -0.01s 0.025%  internal/runtime/maps.typedmemmove
         0     0%   0.2%      0.18s  0.46%  internal/syscall/unix.ARC4Random
         0     0%   0.2%      0.02s 0.051%  internal/syscall/unix.Getaddrinfo
         0     0%   0.2%      0.01s 0.025%  io.ReadAll
         0     0%   0.2%      0.06s  0.15%  io.ReadAtLeast
         0     0%   0.2%      0.17s  0.43%  io.ReadFull (inline)
         0     0%   0.2%     -0.09s  0.23%  net.(*Dialer).DialContext
         0     0%   0.2%     -0.04s   0.1%  net.(*TCPConn).SetKeepAliveConfig
         0     0%   0.2%      0.01s 0.025%  net.(*conn).Close
         0     0%   0.2%     -0.09s  0.23%  net.(*conn).Read
         0     0%   0.2%     -0.01s 0.025%  net.(*conn).SetReadDeadline
         0     0%   0.2%      1.11s  2.81%  net.(*conn).Write
         0     0%   0.2%      0.02s 0.051%  net.(*netFD).Close
         0     0%   0.2%     -0.09s  0.23%  net.(*netFD).Read
         0     0%   0.2%     -0.01s 0.025%  net.(*netFD).SetReadDeadline (inline)
         0     0%   0.2%      1.11s  2.81%  net.(*netFD).Write
         0     0%   0.2%     -0.08s   0.2%  net.(*netFD).connect
         0     0%   0.2%     -0.06s  0.15%  net.(*netFD).dial
         0     0%   0.2%      0.01s 0.025%  net.(*netFD).setAddr (inline)
         0     0%   0.2%     -0.08s   0.2%  net.(*sysDialer).dialParallel
         0     0%   0.2%     -0.10s  0.25%  net.(*sysDialer).dialSingle
         0     0%   0.2%     -0.10s  0.25%  net.(*sysDialer).dialTCP
         0     0%   0.2%     -0.10s  0.25%  net.(*sysDialer).doDialTCP (inline)
         0     0%   0.2%     -0.10s  0.25%  net.(*sysDialer).doDialTCPProto
         0     0%   0.2%      0.02s 0.051%  net._C_getaddrinfo (inline)
         0     0%   0.2%      0.02s 0.051%  net.cgoLookupHostIP
         0     0%   0.2%      0.02s 0.051%  net.cgoLookupIP.func1
         0     0%   0.2%      0.02s 0.051%  net.doBlockingWithCtx[go.shape.[]net.IPAddr].func1
         0     0%   0.2%     -0.07s  0.18%  net.internetSocket
         0     0%   0.2%     -0.03s 0.076%  net.newTCPConn
         0     0%   0.2%      0.03s 0.076%  net.setDefaultSockopts
         0     0%   0.2%     -0.01s 0.025%  net.setKeepAlive
         0     0%   0.2%     -0.02s 0.051%  net.setKeepAliveCount
         0     0%   0.2%      0.01s 0.025%  net.setKeepAliveIdle
         0     0%   0.2%     -0.02s 0.051%  net.setKeepAliveInterval
         0     0%   0.2%      0.01s 0.025%  net.setNoDelay
         0     0%   0.2%     -0.06s  0.15%  net.socket
         0     0%   0.2%     -0.04s   0.1%  net.sysSocket
         0     0%   0.2%      0.01s 0.025%  net/http.(*Cookie).String
         0     0%   0.2%      0.01s 0.025%  net/http.(*Request).Cookie (inline)
         0     0%   0.2%      0.06s  0.15%  net/http.(*conn).readRequest
         0     0%   0.2%      0.24s  0.61%  net/http.(*conn).serve
         0     0%   0.2%      0.10s  0.25%  net/http.(*connReader).Read
         0     0%   0.2%     -0.01s 0.025%  net/http.(*connReader).abortPendingRead
         0     0%   0.2%     -0.08s   0.2%  net/http.(*connReader).backgroundRead
         0     0%   0.2%     -0.01s 0.025%  net/http.(*response).WriteHeader
         0     0%   0.2%      0.21s  0.53%  net/http.(*response).finishRequest
         0     0%   0.2%     -0.01s 0.025%  net/http.Header.Clone (inline)
         0     0%   0.2%     -0.02s 0.051%  net/http.Header.Set (inline)
         0     0%   0.2%     -0.01s 0.025%  net/http.Header.WriteSubset (inline)
         0     0%   0.2%      0.01s 0.025%  net/http.Header.has (inline)
         0     0%   0.2%     -0.01s 0.025%  net/http.Header.sortedKeyValues
         0     0%   0.2%      0.01s 0.025%  net/http.Header.sortedKeyValues.func1
         0     0%   0.2%     -0.01s 0.025%  net/http.Header.writeSubset
         0     0%   0.2%      0.01s 0.025%  net/http.SetCookie
         0     0%   0.2%      0.21s  0.53%  net/http.checkConnErrorWriter.Write
         0     0%   0.2%      0.01s 0.025%  net/http.readCookies
         0     0%   0.2%      0.02s 0.051%  net/http.readRequest
         0     0%   0.2%      0.01s 0.025%  net/http.readTransfer
         0     0%   0.2%      0.01s 0.025%  net/http.sanitizeCookieValue
         0     0%   0.2%     -0.01s 0.025%  net/http.writeStatusLine
         0     0%   0.2%      0.02s 0.051%  net/textproto.(*Reader).ReadMIMEHeader (inline)
         0     0%   0.2%     -0.01s 0.025%  net/textproto.(*Reader).readContinuedLineSlice
         0     0%   0.2%     -0.01s 0.025%  net/textproto.(*Reader).readLineSlice
         0     0%   0.2%     -0.02s 0.051%  net/textproto.MIMEHeader.Set (inline)
         0     0%   0.2%      0.01s 0.025%  net/textproto.canonicalMIMEHeaderKey
         0     0%   0.2%      0.02s 0.051%  net/textproto.readMIMEHeader
         0     0%   0.2%     -0.01s 0.025%  net/url.(*URL).EscapedPath
         0     0%   0.2%      0.02s 0.051%  net/url.(*URL).Query
         0     0%   0.2%      0.01s 0.025%  net/url.Parse
         0     0%   0.2%      0.02s 0.051%  net/url.ParseQuery (inline)
         0     0%   0.2%     -0.01s 0.025%  net/url.ParseRequestURI
         0     0%   0.2%     -0.01s 0.025%  net/url.escape
         0     0%   0.2%     -1.83s  4.63%  os.(*File).Write
         0     0%   0.2%     -1.83s  4.63%  os.(*File).write (inline)
         0     0%   0.2%      0.09s  0.23%  os.Open (inline)
         0     0%   0.2%      0.09s  0.23%  os.OpenFile
         0     0%   0.2%      0.27s  0.68%  os.Stat
         0     0%   0.2%      0.36s  0.91%  os.ignoringEINTR (inline)
         0     0%   0.2%      0.09s  0.23%  os.open (inline)
         0     0%   0.2%      0.09s  0.23%  os.openFileNolog
         0     0%   0.2%      0.09s  0.23%  os.openFileNolog.func1 (inline)
         0     0%   0.2%      0.27s  0.68%  os.statNolog
         0     0%   0.2%      0.27s  0.68%  os.statNolog.func1 (inline)
         0     0%   0.2%      0.01s 0.025%  reflect.(*rtype).Field
         0     0%   0.2%     -0.01s 0.025%  reflect.(*rtype).ptrTo
         0     0%   0.2%      0.01s 0.025%  reflect.(*structType).Field
         0     0%   0.2%     -0.01s 0.025%  reflect.Value.Addr
         0     0%   0.2%      0.01s 0.025%  reflect.Value.Grow
         0     0%   0.2%      0.01s 0.025%  reflect.Value.Interface (inline)
         0     0%   0.2%     -0.01s 0.025%  reflect.Value.IsZero
         0     0%   0.2%     -0.01s 0.025%  reflect.Value.Len (inline)
         0     0%   0.2%      0.01s 0.025%  reflect.Value.grow
         0     0%   0.2%     -0.01s 0.025%  reflect.Value.lenNonSlice
         0     0%   0.2%      0.01s 0.025%  reflect.ValueOf (inline)
         0     0%   0.2%      0.01s 0.025%  reflect.growslice
         0     0%   0.2%     -0.01s 0.025%  reflect.ptrTo (inline)
         0     0%   0.2%      0.01s 0.025%  reflect.valueInterface
         0     0%   0.2%      0.01s 0.025%  runtime.(*inlineUnwinder).next
         0     0%   0.2%      0.01s 0.025%  runtime.(*inlineUnwinder).resolveInternal (inline)
         0     0%   0.2%     -0.01s 0.025%  runtime.(*lfstack).push
         0     0%   0.2%      0.02s 0.051%  runtime.(*mcache).nextFree
         0     0%   0.2%      0.02s 0.051%  runtime.(*mcache).refill
         0     0%   0.2%      0.04s   0.1%  runtime.(*mcentral).grow
         0     0%   0.2%      0.03s 0.076%  runtime.(*mheap).alloc
         0     0%   0.2%      0.06s  0.15%  runtime.(*mheap).allocSpan
         0     0%   0.2%      0.01s 0.025%  runtime.(*mheap).freeSpan (inline)
         0     0%   0.2%      0.01s 0.025%  runtime.(*mheap).freeSpanLocked
         0     0%   0.2%      0.03s 0.076%  runtime.(*mheap).initSpan
         0     0%   0.2%      0.01s 0.025%  runtime.(*mspan).initHeapBits
         0     0%   0.2%      0.01s 0.025%  runtime.(*mspan).typePointersOfUnchecked
         0     0%   0.2%      0.01s 0.025%  runtime.(*pageAlloc).free
         0     0%   0.2%      0.01s 0.025%  runtime.(*pageAlloc).update
         0     0%   0.2%      0.01s 0.025%  runtime.(*sweepLocked).sweep
         0     0%   0.2%      0.01s 0.025%  runtime.(*sweepLocked).sweep.(*mheap).freeSpan.func2
         0     0%   0.2%      0.36s  0.91%  runtime.(*timer).maybeAdd
         0     0%   0.2%      0.35s  0.88%  runtime.(*timer).modify
         0     0%   0.2%      0.36s  0.91%  runtime.(*timer).reset (inline)
         0     0%   0.2%      0.01s 0.025%  runtime.(*unwinder).next
         0     0%   0.2%     -0.01s 0.025%  runtime.acquirep
         0     0%   0.2%      0.01s 0.025%  runtime.bgsweep
         0     0%   0.2%      0.02s 0.051%  runtime.callers
         0     0%   0.2%      0.02s 0.051%  runtime.callers.func1
         0     0%   0.2%     -0.02s 0.051%  runtime.casgstatus
         0     0%   0.2%     -0.01s 0.025%  runtime.checkRunqsNoP
         0     0%   0.2%     -0.03s 0.076%  runtime.closechan.goready.func1
         0     0%   0.2%      0.01s 0.025%  runtime.convT32
         0     0%   0.2%     -0.02s 0.051%  runtime.convTstring
         0     0%   0.2%     -0.02s 0.051%  runtime.entersyscall_sysmon
         0     0%   0.2%     -0.02s 0.051%  runtime.findObject
         0     0%   0.2%     -0.14s  0.35%  runtime.findRunnable
         0     0%   0.2%     -0.01s 0.025%  runtime.forEachPInternal
         0     0%   0.2%     -0.01s 0.025%  runtime.gcBgMarkWorker
         0     0%   0.2%     -0.01s 0.025%  runtime.gcBgMarkWorker.func1
         0     0%   0.2%      0.03s 0.076%  runtime.gcBgMarkWorker.func2
         0     0%   0.2%      0.03s 0.076%  runtime.gcDrain
         0     0%   0.2%     -0.03s 0.076%  runtime.gcDrainMarkWorkerDedicated (inline)
         0     0%   0.2%      0.06s  0.15%  runtime.gcDrainMarkWorkerIdle (inline)
         0     0%   0.2%     -0.01s 0.025%  runtime.gcMarkTermination.forEachP.func6
         0     0%   0.2%     -0.01s 0.025%  runtime.gcMarkTermination.func3
         0     0%   0.2%     -0.10s  0.25%  runtime.goexit0
         0     0%   0.2%     -0.02s 0.051%  runtime.growslice
         0     0%   0.2%     -0.11s  0.28%  runtime.injectglist
         0     0%   0.2%     -0.01s 0.025%  runtime.lfstackPack (inline)
         0     0%   0.2%     -0.11s  0.28%  runtime.lock (inline)
         0     0%   0.2%     -0.12s   0.3%  runtime.lockWithRank (inline)
         0     0%   0.2%      0.04s   0.1%  runtime.mPark (inline)
         0     0%   0.2%      0.01s 0.025%  runtime.mProf_Malloc
         0     0%   0.2%     -0.01s 0.025%  runtime.mProf_Malloc.func1
         0     0%   0.2%      0.01s 0.025%  runtime.makechan
         0     0%   0.2%      0.03s 0.076%  runtime.makemap
         0     0%   0.2%      0.01s 0.025%  runtime.makemap_small
         0     0%   0.2%      0.01s 0.025%  runtime.makeslice
         0     0%   0.2%      0.05s  0.13%  runtime.mallocgcSmallScanHeader
         0     0%   0.2%     -0.01s 0.025%  runtime.mapIterStart
         0     0%   0.2%      0.03s 0.076%  runtime.mapaccess1_faststr
         0     0%   0.2%      0.01s 0.025%  runtime.mapaccess2_faststr
         0     0%   0.2%     -0.02s 0.051%  runtime.mapassign
         0     0%   0.2%      0.01s 0.025%  runtime.mapdelete
         0     0%   0.2%      0.01s 0.025%  runtime.mapdelete_faststr
         0     0%   0.2%     -0.02s 0.051%  runtime.markroot
         0     0%   0.2%     -0.02s 0.051%  runtime.markroot.func1
         0     0%   0.2%     -0.23s  0.58%  runtime.mcall
         0     0%   0.2%     -0.01s 0.025%  runtime.morestack
         0     0%   0.2%     -0.02s 0.051%  runtime.nanotime (inline)
         0     0%   0.2%      0.01s 0.025%  runtime.netpollBreak (inline)
         0     0%   0.2%     -0.01s 0.025%  runtime.netpollblock
         0     0%   0.2%      0.04s   0.1%  runtime.netpollgoready.goready.func1
         0     0%   0.2%     -0.02s 0.051%  runtime.netpollopen
         0     0%   0.2%     -0.01s 0.025%  runtime.netpollready
         0     0%   0.2%      0.07s  0.18%  runtime.newarray
         0     0%   0.2%     -0.05s  0.13%  runtime.newobject
         0     0%   0.2%      0.01s 0.025%  runtime.newproc
         0     0%   0.2%      0.12s   0.3%  runtime.newproc.func1
         0     0%   0.2%      0.01s 0.025%  runtime.newproc1
         0     0%   0.2%     -0.01s 0.025%  runtime.newstack
         0     0%   0.2%      0.04s   0.1%  runtime.notesleep
         0     0%   0.2%      0.10s  0.25%  runtime.notewakeup
         0     0%   0.2%     -0.13s  0.33%  runtime.osyield (inline)
         0     0%   0.2%     -0.13s  0.33%  runtime.park_m
         0     0%   0.2%      0.01s 0.025%  runtime.pcdatavalue1
         0     0%   0.2%      0.01s 0.025%  runtime.pcvalue
         0     0%   0.2%      0.05s  0.13%  runtime.pollWork
         0     0%   0.2%     -0.01s 0.025%  runtime.preemptM
         0     0%   0.2%     -0.01s 0.025%  runtime.preemptPark
         0     0%   0.2%      0.01s 0.025%  runtime.processWakeupEvent (inline)
         0     0%   0.2%      0.01s 0.025%  runtime.profilealloc
         0     0%   0.2%     -0.01s 0.025%  runtime.rawbyteslice
         0     0%   0.2%     -0.05s  0.13%  runtime.ready
         0     0%   0.2%     -0.06s  0.15%  runtime.resetspinning
         0     0%   0.2%      0.01s 0.025%  runtime.runqgrab
         0     0%   0.2%      0.01s 0.025%  runtime.runqsteal
         0     0%   0.2%     -0.01s 0.025%  runtime.scanstack
         0     0%   0.2%     -0.01s 0.025%  runtime.semacquire1
         0     0%   0.2%      0.04s   0.1%  runtime.semasleep
         0     0%   0.2%      0.11s  0.28%  runtime.semawakeup
         0     0%   0.2%     -0.01s 0.025%  runtime.semrelease (inline)
         0     0%   0.2%     -0.01s 0.025%  runtime.semrelease1
         0     0%   0.2%     -0.06s  0.15%  runtime.send.goready.func1
         0     0%   0.2%     -0.01s 0.025%  runtime.signalM (inline)
         0     0%   0.2%      0.02s 0.051%  runtime.slicebytetostring
         0     0%   0.2%     -0.01s 0.025%  runtime.startTheWorldWithSema
         0     0%   0.2%      0.11s  0.28%  runtime.startm
         0     0%   0.2%     -0.02s 0.051%  runtime.stopm
         0     0%   0.2%     -0.01s 0.025%  runtime.suspendG
         0     0%   0.2%      0.01s 0.025%  runtime.sweepone
         0     0%   0.2%      0.02s 0.051%  runtime.sysUsed (inline)
         0     0%   0.2%      0.02s 0.051%  runtime.sysUsedOS (inline)
         0     0%   0.2%      0.02s 0.051%  runtime.tracebackPCs
         0     0%   0.2%      0.01s 0.025%  runtime.unlock2Wake
         0     0%   0.2%      0.34s  0.86%  runtime.wakeNetPoller
         0     0%   0.2%      0.02s 0.051%  runtime.wakeNetpoll
         0     0%   0.2%      0.22s  0.56%  runtime.wakep
         0     0%   0.2%     -0.01s 0.025%  slices.SortFunc[go.shape.[]net/http.keyValues,go.shape.struct { net/http.key string; net/http.values []string }] (inline)
         0     0%   0.2%      0.03s 0.076%  sort.Sort
         0     0%   0.2%      0.01s 0.025%  sort.median
         0     0%   0.2%      0.01s 0.025%  sort.medianAdjacent (inline)
         0     0%   0.2%      0.01s 0.025%  sort.order2 (inline)
         0     0%   0.2%      0.01s 0.025%  strconv.AppendInt
         0     0%   0.2%     -0.01s 0.025%  strconv.AppendQuote (inline)
         0     0%   0.2%      0.01s 0.025%  strconv.FormatFloat (inline)
         0     0%   0.2%     -0.01s 0.025%  strconv.ParseFloat
         0     0%   0.2%     -0.01s 0.025%  strconv.appendEscapedRune
         0     0%   0.2%     -0.01s 0.025%  strconv.appendQuotedWith
         0     0%   0.2%     -0.01s 0.025%  strconv.atof64
         0     0%   0.2%     -0.01s 0.025%  strconv.parseFloatPrefix
         0     0%   0.2%      0.01s 0.025%  strings.(*Builder).WriteString (inline)
         0     0%   0.2%     -0.01s 0.025%  strings.(*Replacer).Replace
         0     0%   0.2%      0.01s 0.025%  strings.Compare (inline)
         0     0%   0.2%      0.01s 0.025%  strings.Join
         0     0%   0.2%      0.01s 0.025%  sync.(*Map).Load (inline)
         0     0%   0.2%      0.01s 0.025%  sync.(*Mutex).Lock (inline)
         0     0%   0.2%     -0.01s 0.025%  sync.(*Once).Do (inline)
         0     0%   0.2%      0.01s 0.025%  sync.(*Pool).Get
         0     0%   0.2%     -0.01s 0.025%  sync.(*poolChain).pushHead
         0     0%   0.2%      0.02s 0.051%  syscall.Close
         0     0%   0.2%      0.01s 0.025%  syscall.CloseOnExec (inline)
         0     0%   0.2%     -0.02s 0.051%  syscall.Connect
         0     0%   0.2%     -0.02s 0.051%  syscall.Getpeername
         0     0%   0.2%      0.01s 0.025%  syscall.Getsockname
         0     0%   0.2%     -0.02s 0.051%  syscall.GetsockoptInt
         0     0%   0.2%      0.09s  0.23%  syscall.Open
         0     0%   0.2%     -0.08s   0.2%  syscall.Read (inline)
         0     0%   0.2%      0.04s   0.1%  syscall.SetNonblock
         0     0%   0.2%     -0.09s  0.23%  syscall.Socket
         0     0%   0.2%      0.27s  0.68%  syscall.Stat
         0     0%   0.2%     -0.70s  1.77%  syscall.Write (inline)
         0     0%   0.2%     -0.02s 0.051%  syscall.connect
         0     0%   0.2%      0.05s  0.13%  syscall.fcntl
         0     0%   0.2%     -0.02s 0.051%  syscall.getpeername
         0     0%   0.2%      0.01s 0.025%  syscall.getsockname
         0     0%   0.2%     -0.02s 0.051%  syscall.getsockopt
         0     0%   0.2%     -0.08s   0.2%  syscall.read
         0     0%   0.2%     -0.09s  0.23%  syscall.socket
         0     0%   0.2%     -0.70s  1.77%  syscall.write
         0     0%   0.2%      0.37s  0.94%  time.(*Timer).Reset
         0     0%   0.2%     -0.01s 0.025%  time.AfterFunc
         0     0%   0.2%      0.01s 0.025%  time.Now
         0     0%   0.2%     -0.01s 0.025%  time.Until
         0     0%   0.2%     -0.01s 0.025%  time.newTimer
         0     0%   0.2%      0.36s  0.91%  time.resetTimer
         0     0%   0.2%      0.01s 0.025%  time.runtimeNano
         0     0%   0.2%      0.01s 0.025%  time.runtimeNow
         0     0%   0.2%      0.01s 0.025%  time.when