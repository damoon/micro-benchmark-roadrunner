<?php
/**
 * Dead simple, high performance, drop-in bridge to Golang RPC with zero dependencies
 *
 * @author Wolfy-J
 */

namespace Spiral\Goridge;

/**
 * Communicates with remote server/client over streams using byte payload:
 *
 * [ prefix       ][ payload                               ]
 * [ 1+8+8 bytes  ][ message length|LE ][message length|BE ]
 *
 * prefix:
 * [ flag       ][ message length, unsigned int 64bits, LittleEndian ]
 */
class StreamRelay implements RelayInterface
{
    /** @var resource */
    private $in;

    /** @var resource */
    private $out;

    /**
     * Example:
     * $relay = new StreamRelay(STDIN, STDOUT);
     *
     * @param resource $in  Must be readable.
     * @param resource $out Must be writable.
     *
     * @throws Exceptions\InvalidArgumentException
     */
    public function __construct($in, $out)
    {
        if (!$this->assertMode($in, 'r')) {
            throw new Exceptions\InvalidArgumentException("resource `in` must be readable");
        }

        if (!$this->assertMode($out, 'w')) {
            throw new Exceptions\InvalidArgumentException("resource `out` must be readable");
        }

        $this->in = $in;
        $this->out = $out;
    }

    /**
     * {@inheritdoc}
     */
    public function send($payload, int $flags = null)
    {
        $size = strlen($payload);
        if ($flags & self::PAYLOAD_NONE && $size != 0) {
            throw new Exceptions\TransportException("unable to send payload with PAYLOAD_NONE flag");
        }

        if (fwrite($this->out, pack('CPJ', $flags, $size, $size), 17) === false) {
            throw new Exceptions\TransportException("unable to write prefix to the stream");
        }

        if (!($flags & self::PAYLOAD_NONE) && fwrite($this->out, $payload, $size) === false) {
            throw new Exceptions\TransportException("unable to write payload to the stream");
        }

        return $this;
    }

    /**
     * {@inheritdoc}
     */
    public function receiveSync(int &$flags = null)
    {
        $prefix = $this->fetchPrefix();
        $flags = $prefix['flags'];

        $result = null;
        if ($prefix['size'] !== 0) {
            $leftBytes = $prefix['size'];

            //Add ability to write to stream in a future
            while ($leftBytes > 0) {
                $buffer = fread($this->in, min($leftBytes, self::BUFFER_SIZE));
                if ($buffer === false) {
                    throw new Exceptions\TransportException("error reading payload from the stream");
                }

                $result .= $buffer;
                $leftBytes -= strlen($buffer);
            }
        }

        return $result;
    }

    /**
     * @return array Prefix [flag, length]
     *
     * @throws Exceptions\PrefixException
     */
    private function fetchPrefix(): array
    {
        $prefixBody = fread($this->in, 17);
        if ($prefixBody === false) {
            throw new Exceptions\PrefixException("unable to read prefix from the stream");
        }

        $result = unpack("Cflags/Psize/Jrevs", $prefixBody);
        if (!is_array($result)) {
            throw new Exceptions\PrefixException("invalid prefix");
        }

        if ($result['size'] != $result['revs']) {
            throw new Exceptions\PrefixException("invalid prefix (checksum)");
        }

        return $result;
    }

    /**
     * Checks if stream is writable or readable.
     *
     * @param resource $stream
     * @param string   $mode
     *
     * @return bool
     */
    private function assertMode($stream, $mode): bool
    {
        $meta = stream_get_meta_data($stream);

        return strpos($meta['mode'], $mode) !== false;
    }
}
