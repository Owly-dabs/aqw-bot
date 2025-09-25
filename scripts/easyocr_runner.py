import easyocr
import sys
import json

def main():
    image_path = sys.argv[1] if len(sys.argv) > 1 else None
    if not image_path:
        print(json.dumps({"error": "No image path provided"}))
        return

    reader = easyocr.Reader(['en'], gpu=False)
    results = reader.readtext(image_path)

    # Extract just the text (you can include confidence if needed)
    output = [text for (_, text, _) in results]
    print(json.dumps({"text": output}))

if __name__ == "__main__":
    main()